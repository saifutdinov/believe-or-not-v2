package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379", // или localhost:6379 вне Docker
	})
}
