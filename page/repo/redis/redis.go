package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/unknowntpo/page/domain"

	"github.com/unknowntpo/page/pkg/errors"

	"github.com/redis/go-redis/v9"
)

type pageRepoImpl struct {
	// any fields needed for implementation
	client *redis.Client
}

func NewPageRepo(c *redis.Client) domain.PageRepo {
	return &pageRepoImpl{client: c}
}

// NewList initializes the list with listMeta, pageList.
// If list already exists, we return domain.ErrListAlreadyExists
func (r *pageRepoImpl) NewList(ctx context.Context, userID int64, listKey domain.ListKey) error {
	listMetaKeyByUser := domain.GenerateListMetaKeyByUserID(listKey, userID)
	keys := []string{string(listMetaKeyByUser)}

	// Create a Lua script to get the max score and add a new value
	script := redis.NewScript(fmt.Sprintf(`
		redis.log(redis.LOG_NOTICE, "got KEYS", KEYS[1])

		-- if listMeta exist, return error
		if redis.call("EXISTS", KEYS[1]) == 1 then
			return {err = "%s"}
		end
		-- init listMeta, set head, tail to ""
		redis.call("HSET", KEYS[1], "head", "", "tail", "")

		return {ok = "status"}
	`, domain.ErrListAlreadyExists.Error()))

	_, err := script.Run(context.Background(), r.client, keys).Result()
	if err != nil {
		switch {
		case strings.Contains(err.Error(), domain.ErrListAlreadyExists.Error()):
			return domain.ErrListAlreadyExists
		default:
			return errors.Wrap(errors.Internal, " failed on script.Run", err)
		}
	}
	return nil
}

// GetPage gets the page by `pageKey`.
// If page not found, we return domain.ErrPageNotFound
func (r *pageRepoImpl) GetPage(ctx context.Context, _ int64, _ domain.ListKey, pageKey domain.PageKey) (domain.Page, error) {
	keys := []string{string(pageKey)}
	args := []any{}

	// Create a Lua script to get the max score and add a new value
	script := redis.NewScript(fmt.Sprintf(`
		-- get all content
		local page = redis.call("JSON.Get", KEYS[1], '.')
		if not page then
			return {err = '%s'}
		end
		return page
	`, domain.ErrPageNotFound))

	result, err := script.Run(context.Background(), r.client, keys, args...).Result()
	if err != nil {
		switch {
		case strings.Contains(err.Error(), domain.ErrPageNotFound.Error()):
			return domain.Page{}, domain.ErrPageNotFound
		}
		return domain.Page{}, errors.Wrap(errors.Internal, " failed on script.Run", err)
	}

	p := domain.Page{}
	if err := json.Unmarshal([]byte(result.(string)), &p); err != nil {
		return domain.Page{}, errors.Wrap(errors.Internal, "failed on json.Unmarshal", err)
	}
	return p, nil
}

// GetHead gets the head for head field of HashMap listMeta, but if head inside it is expired,
// oldest valid pageKey inside sortedset `pageList` will be returned, and as a side effect,
// expired pageKey will be cleaned.
// If listKey not exist, we return domain.ErrListNotFound.
func (r *pageRepoImpl) GetHead(ctx context.Context, userID int64, listKey domain.ListKey) (domain.PageKey, error) {
	listKeyByUser := domain.GenerateListKeyByUserID(listKey, userID)
	listMetaKey := domain.GenerateListMetaKeyByUserID(listKey, userID)

	keys := []string{string(listMetaKey), string(listKeyByUser)}
	args := []any{
		time.Now().Add(domain.DefaultPageTTL).UnixNano(),
	}

	// Create a Lua script to get the max score and add a new value
	script := redis.NewScript(fmt.Sprintf(`
    local listMetaKey = KEYS[1]
    local listKeyByUser = KEYS[2]
    local expireTime = ARGV[1]

	if redis.call("EXISTS", listMetaKey) == 0 then
      return { err = "%s" }
    end

	-- get head from listMeta
	local headPageKey = redis.call("HGET", listMetaKey, "head")

    -- edge case: list exist but list has no page
    if headPageKey == "" then
      return ""
    end

	-- check if head does exist
	if redis.call("EXISTS", headPageKey) == 0 then
		-- means head is expired, use ZREMRANGEBYSCORE to remove expired key
		redis.call("ZREMRANGEBYSCORE", listKeyByUser, 0, expireTime)

		-- -- set pageMeta.head to oldest key that doesn't expired
		headPageKey = redis.call("ZRANGE", KEYS[2], 0, "+inf", "BYSCORE", "LIMIT", 0, 1)
		redis.log(redis.LOG_NOTICE, "got headPageKey", headPageKey)
		redis.call("HSET", KEYS[1], "head", headPageKey)

		-- return the new one
		return headPageKey
	end

	return headPageKey
	`, domain.ErrListNotFound.Error()))

	result, err := script.Run(context.Background(), r.client, keys, args...).Result()
	if err != nil {
		switch {
		case strings.Contains(err.Error(), domain.ErrListNotFound.Error()):
			return "", domain.ErrListNotFound
		default:
			return "", errors.Wrap(errors.Internal, " failed on script.Run", err)
		}
	}
	fmt.Println("\ngot pageHead", result.(string))
	return domain.PageKey(result.(string)), nil
}

// SetPage do the following things:
// - Find the tail recorded in listMeta.tail, modify .next field of tail RedisJSON object point to itself.
// - Add new pageKey with expire time to sorted set (pageList)
// - Add new RedisJSON object
// If listKey not exist, we return domain.ErrListNotFound.
func (r *pageRepoImpl) SetPage(
	ctx context.Context,
	userID int64,
	listKey domain.ListKey,
	p domain.Page,
) (domain.PageKey, error) {
	return r.setPage(ctx, userID, listKey, p)
}

func (r *pageRepoImpl) setPage(
	ctx context.Context,
	userID int64,
	listKey domain.ListKey,
	p domain.Page,
) (domain.PageKey, error) {
	// implementation
	listKeyByUser := domain.GenerateListKeyByUserID(listKey, userID)
	listMetaKeyByUser := domain.GenerateListMetaKeyByUserID(listKey, userID)

	// NOTE: the time.Time to generate pageKey should be the same as the one used
	// to set score (expired time of the pageKey in sortedset).
	// Otherwise, the order of expired time can't be ensured
	now := time.Now()
	p.Key = domain.GeneratePageKeyByListKeyUserID(listKey, userID, now)

	pageContent := p.Marshal()

	keys := []string{
		string(listMetaKeyByUser),
		string(listKeyByUser),
		string(p.Key),
	}
	args := []any{
		// actual page data
		pageContent,
		// score of the page we wanna add (will be expire time of pageKey)
		// NOTE: Use UnixNano to avoid collision on same score
		now.Add(domain.DefaultPageTTL).UnixNano(),
	}

	ttl := int(domain.DefaultPageTTL.Seconds())

	// If listMeta doesn't exist, return error
	// If there's no element in list
	script := redis.NewScript(fmt.Sprintf(`
		local listMetaKeyByUser = KEYS[1]
		local listKeyByUser = KEYS[2]
		local pageKey = KEYS[3]
		local pageContent = ARGV[1]
		local dueTime = ARGV[2]

		if redis.call("EXISTS", listMetaKeyByUser) == 0 then
			-- HashMap doesn't exist, return error
			return {err = "%s"}
		end

		if redis.call("ZADD",listKeyByUser, dueTime, pageKey) ~= 1 then
			return {err = "failed to add pageKey to sorted set"}
		end

		-- Set listMeta.head = pageKey if there's no element in list (head == "")
		if redis.call("HGET", listMetaKeyByUser, "head") == "" then
			redis.call("HSET", listMetaKeyByUser, "head", pageKey)
		else
			-- get old listMeta.tail
			local oldTailPageKey = redis.call("HGET", listMetaKeyByUser, "tail")
			-- local quotedPageKey = "'" .. pageKey .. "'"
			local quotedPageKey = [["]] .. pageKey .. [["]]
			redis.call("JSON.SET", oldTailPageKey, ".next", quotedPageKey)
		end

		-- Set pageMeta.tail = pageKey
		redis.call("HSET", listMetaKeyByUser, "tail", pageKey)

		-- set key: pageKey to actual data with 1 day TTL
		redis.call('JSON.SET', pageKey, '.', pageContent)
		redis.call('EXPIRE', pageKey, %d)

		return pageKey
	`, domain.ErrListNotFound.Error(), ttl))

	result, err := script.Run(context.Background(), r.client, keys, args...).Result()
	if err != nil {
		switch {
		case strings.Contains(err.Error(), domain.ErrListNotFound.Error()):
			return "", domain.ErrListNotFound
		default:
			return "", errors.Wrap(errors.Internal, " failed on script.Run", err)
		}
	}

	pageKey := domain.PageKey(result.(string))

	return pageKey, nil
}
