package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/unknowntpo/page/internal/domain"

	"github.com/redis/go-redis/v9"
)

type pageRepoImpl struct {
	// any fields needed for implementation
	client *redis.Client
}

func NewPageRepo(c *redis.Client) domain.PageRepo {
	return &pageRepoImpl{client: c}
}

func (r *pageRepoImpl) GetPage(ctx context.Context, pageKey domain.PageKey) (domain.Page, error) {
	// implementation

	// if page doesn't exist, return PageDoesNotExist error
	return domain.Page{}, nil
}

func (r *pageRepoImpl) GetHead(ctx context.Context, listKey domain.ListKey) (domain.PageKey, error) {
	// implementation

	// return the first element of the page range from now() - 1 Day to now()
	return domain.PageKey("asdf"), nil
}

func (r *pageRepoImpl) SetPage(
	ctx context.Context,
	userID int64,
	listKey domain.ListKey,
	p domain.Page,
) error {
	return r.setPage(ctx, userID, listKey, p)
}

func (r *pageRepoImpl) setPage(
	ctx context.Context,
	userID int64,
	listKey domain.ListKey,
	p domain.Page,
) error {
	// implementation
	listKeyByUser := domain.GenerateListKeyByUserID(listKey, userID)
	headIfHashMapNotExist := domain.GeneratePageKey()
	nextCandidate := domain.GeneratePageKey()
	pageMetaKey := domain.GeneratePageMetaKeyByUserID(listKey, userID)

	// Create a Lua script to get the max score and add a new value
	script := redis.NewScript(`
		redis.log(redis.LOG_NOTICE, "got KEYS", KEYS[1], KEYS[2])
		redis.log(redis.LOG_NOTICE, "got ARGV", ARGV[1], ARGV[2], ARGV[3])

		-- Add pageMeta.nextCandidate to sorted set pageList with score: ARGV[2]
		local pageKey
		if redis.call("EXISTS", KEYS[1]) == 0 then
			-- HashMap doesn't exist, create new one 
			redis.call("HSET", KEYS[1], "head", ARGV[4])
			redis.call("HSET", KEYS[1], "tail", ARGV[4])
			redis.call("HSET", KEYS[1], "nextCandidate", ARGV[3])
			pageKey = ARGV[3]
		else
			pageKey = redis.call("HGET", KEYS[1], "nextCandidate")
			redis.log(redis.LOG_NOTICE, "got pageKey and score", pageKey, ARGV[2])
		end
		local res = redis.call("ZADD", KEYS[2], ARGV[2], pageKey)
		if res ~= 1 then
			return {err = string.format("failed to set new value to pageList with pageKey: %s, score", pageKey, ARGV[2])}
		end

		-- set key: KEYS[1] to ARGV[1]
		local res = redis.call("SET", KEYS[2], ARGV[1])
		if res ~= 1 then
			return {err = string.format("failed to set page with key: %s", KEY[2])}
		end
		-- Set pageMeta.nextCandidate = ARGV[3] (new candidate)
		res = redis.call("HSET", KEYS[1], "nextCandidate", ARGV[3])
		if res ~= 1 then
			return {err = string.format("failed to set pageMeta.nextCandidate with key: %s", ARGV[3])}
		end
		-- Set pageMeta.tail = pageKey
		res = redis.call("HSET", KEYS[1], "tail", pageKey)
		if res ~= 1 then
			return {err = string.format("failed to set pageMeta.tail with key: %s", pageKey)}
		end

		redis.log(redis.LOG_NOTICE, "doneWithValue", ARGV[1])
	`)

	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	keys := []string{string(pageMetaKey), string(listKeyByUser)}
	args := []any{
		string(b),
		time.Now().Unix(),
		string(nextCandidate),
		string(headIfHashMapNotExist),
	}

	result, err := script.Run(context.Background(), r.client, keys, args...).Result()
	if err != nil {
		return err
	}

	// Print the result
	fmt.Printf("Result: %v", result)

	return nil
}
