package redis

import (
	"DynamicLED/config"
	"github.com/go-redis/redis"
)

func NewClient(config *config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DB,
	})
}
