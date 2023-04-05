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

func (r *pageRepoImpl) NewList(ctx context.Context, userID int64, listKey domain.ListKey) error {
	listMetaKeyByUser := domain.GenerateListMetaKeyByUserID(listKey, userID)

	// Create a Lua script to get the max score and add a new value
	script := redis.NewScript(`
		redis.log(redis.LOG_NOTICE, "got KEYS", KEYS[1])

		-- KEYS[1]: listMetaKeyByUser

		-- if listMeta exist, return error
		if redis.call("EXISTS", KEYS[1]) == 1 then
			return {err = "list has already exist"}
		end
		-- init listMeta, set head, tail, nextCandidate to ""
		redis.call("HSET", KEYS[1], "head", "", "tail", "", "nextCandidate", "")

		return {ok = "status"}
	`)

	keys := []string{string(listMetaKeyByUser)}

	_, err := script.Run(context.Background(), r.client, keys).Result()
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "ResourceAlreadyExist"):
			return errors.Wrap(errors.ResourceAlreadyExist, fmt.Sprintf("pageList %s for userID [%d] has already exist", listKey, userID), err)
		default:
			return errors.Wrap(errors.Internal, " failed on script.Run", err)
		}
	}
	return nil
}

func (r *pageRepoImpl) GetPage(ctx context.Context, pageKey domain.PageKey) (domain.Page, error) {
	// implementation
	pageStr, err := r.client.Get(ctx, string(pageKey)).Result()
	if err != nil {
		switch err {
		case redis.Nil:
			return domain.Page{}, errors.Wrap(errors.ResourceNotFound, "", err)
		default:
			return domain.Page{}, errors.Wrap(errors.Internal, "failed on r.client.Get", err)
		}
	}
	p := domain.Page{}
	if err := json.Unmarshal([]byte(pageStr), &p); err != nil {
		return domain.Page{}, errors.Wrap(errors.Internal, "failed on json.Unmarshal", err)
	}
	return p, nil
}

func (r *pageRepoImpl) GetHead(ctx context.Context, userID int64, listKey domain.ListKey) (domain.PageKey, error) {
	listKeyByUser := domain.GenerateListKeyByUserID(listKey, userID)
	listMetaKey := domain.GenerateListMetaKeyByUserID(listKey, userID)

	// Create a Lua script to get the max score and add a new value
	script := redis.NewScript(`
		redis.log(redis.LOG_NOTICE, "got KEYS", KEYS[1], KEYS[2])
		redis.log(redis.LOG_NOTICE, "got ARGV", ARGV[1])

		-- KEYS[1]: pageMetaKey
		-- KEYS[2]: listKeyByUser
		-- ARGV[1]: currentTimestamp

		-- get head from pageMeta
		local headPageKey = redis.call("HGET", KEYS[1], "head")
		redis.log(redis.LOG_NOTICE, "before check, got headPageKey", headPageKey)

		-- check if head does exist
		if redis.call("EXISTS", headPageKey) == 0 then
			-- not exist
			-- -- means head is expired, use ZREMRANGEBYSCORE to remove expired key
			redis.call("ZREMRANGEBYSCORE", KEYS[2], "-inf", ARGV[1])
			-- -- set pageMeta.head to oldest key that doesn't expired
			headPageKey = redis.call("ZRANGE", KEYS[2], 0, "+inf", "BYSCORE", "LIMIT", 0, 1)
			redis.log(redis.LOG_NOTICE, "got headPageKey", headPageKey)
			redis.call("HSET", KEYS[1], "head", headPageKey)
		end

		redis.log(redis.LOG_NOTICE, "after check, got headPageKey", headPageKey)

		return headPageKey
	`)

	keys := []string{string(listMetaKey), string(listKeyByUser)}
	args := []any{
		time.Now().Unix(),
	}

	result, err := script.Run(context.Background(), r.client, keys, args...).Result()
	if err != nil {
		return "", errors.Wrap(errors.Internal, " failed on script.Run", err)
	}

	fmt.Println("\ngot pageHead", result.(string))
	return domain.PageKey(result.(string)), nil
}

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
	headIfHashMapNotExist := domain.GeneratePageKey()
	nextCandidate := domain.GeneratePageKey()
	pageMetaKey := domain.GenerateListMetaKeyByUserID(listKey, userID)

	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	keys := []string{string(pageMetaKey), string(listKeyByUser)}
	args := []any{
		string(b),
		// score will be expire time of pageKey
		time.Now().Add(24 * time.Hour).Unix(),
		string(nextCandidate),
		string(headIfHashMapNotExist),
	}

	// Create a Lua script to get the max score and add a new value
	script := redis.NewScript(`
		redis.log(redis.LOG_NOTICE, "got KEYS", KEYS[1], KEYS[2])
		redis.log(redis.LOG_NOTICE, "got ARGV", ARGV[1], ARGV[2], ARGV[3])

		-- Add pageMeta.nextCandidate to sorted set pageList with score: ARGV[2]
		local pageKey
		if redis.call("EXISTS", KEYS[1]) == 0 then
			-- HashMap doesn't exist, return error
			return {err = "ResourceNotFound"}
		else
			pageKey = redis.call("HGET", KEYS[1], "nextCandidate")
			redis.log(redis.LOG_NOTICE, "got pageKey and score", pageKey, ARGV[2])
		end
		local res = redis.call("ZADD", KEYS[2], ARGV[2], pageKey)

		-- set key: KEYS[1] to ARGV[1] with 1 day TTL
		redis.call("SET", ARGV[4], ARGV[1], "EX", "86400")

		-- Set pageMeta.nextCandidate = ARGV[3] (new candidate)
		redis.call("HSET", KEYS[1], "nextCandidate", ARGV[3])
		-- Set pageMeta.tail = pageKey
		redis.call("HSET", KEYS[1], "tail", pageKey)

		redis.log(redis.LOG_NOTICE, "doneWithValue", ARGV[1])

		return pageKey
	`)

	result, err := script.Run(context.Background(), r.client, keys, args...).Result()
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "ResourceNotFound"):
			return "", errors.Wrap(errors.ResourceNotFound, fmt.Sprintf("pageList %s for userID [%d] not found, call NewList first", listKey, userID), err)
		default:
			return "", errors.Wrap(errors.Internal, " failed on script.Run", err)
		}
	}

	pageKey := domain.PageKey(result.(string))

	return pageKey, nil
}
