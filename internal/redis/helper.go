package redis

import (
	"context"

	"github.com/unknowntpo/page/internal/infra"

	redis "github.com/redis/go-redis/v9"
)

func PrepareTestDatabase() *redis.Client {
	client := infra.NewRedisClient()
	client.FlushAll(context.Background())
	return client
}
