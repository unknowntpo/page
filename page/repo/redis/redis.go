package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/unknowntpo/page/domain"

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

func (r *pageRepoImpl) GetHead(ctx context.Context, userID int64, listKey domain.ListKey) (domain.PageKey, error) {
	listKeyByUser := domain.GenerateListKeyByUserID(listKey, userID)
	pageMetaKey := domain.GeneratePageMetaKeyByUserID(listKey, userID)

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

		return headPageKey
	`)

	keys := []string{string(pageMetaKey), string(listKeyByUser)}
	args := []any{
		time.Now().Unix(),
	}

	result, err := script.Run(context.Background(), r.client, keys, args...).Result()
	if err != nil {
		return "", err
	}

	fmt.Println("got pageHead", result)
	return domain.PageKey(result.(string)), nil
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

		-- set key: KEYS[1] to ARGV[1] with 1 day TTL
		redis.call("SET", ARGV[4], ARGV[1], "EX", "86400")

		-- Set pageMeta.nextCandidate = ARGV[3] (new candidate)
		redis.call("HSET", KEYS[1], "nextCandidate", ARGV[3])
		-- Set pageMeta.tail = pageKey
		redis.call("HSET", KEYS[1], "tail", pageKey)

		redis.log(redis.LOG_NOTICE, "doneWithValue", ARGV[1])

		return {ok = "status"}	
	`)

	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	keys := []string{string(pageMetaKey), string(listKeyByUser)}
	args := []any{
		string(b),
		// score will be expire time of pageKey
		time.Now().Add(24 * time.Hour).Unix(),
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
