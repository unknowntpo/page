package redis

import (
	"github.com/unknowntpo/page/pkg/redis"

	goRedis "github.com/redis/go-redis/v9"
)

func PrepareTestDatabase() *goRedis.Client {
	return redis.NewClient()
}
