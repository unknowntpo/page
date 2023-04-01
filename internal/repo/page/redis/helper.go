package redis

import (
	"github.com/unknowntpo/page/internal/infra"

	redis "github.com/redis/go-redis/v9"
)

func PrepareTestDatabase() *redis.Client {
	return infra.NewRedisClient()
}
