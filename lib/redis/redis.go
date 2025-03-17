package redis

import (
	"my-project-be/config"

	"github.com/redis/go-redis/v9"
)

func RedisClient(cfg *config.AppConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Username: "default",
		Password: cfg.RedisPass,
		DB:       0,
	})

	return rdb
}

