package services

import (
	"fmt"
	"github.com/ditojanelidze/jwt-auth/config/initializers"
	"github.com/go-redis/redis/v7"
)

func redisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", initializers.REDIS_URL, initializers.REDIS_PORT),
		Password: "",
		DB:       initializers.REDIS_DB,
	})
}