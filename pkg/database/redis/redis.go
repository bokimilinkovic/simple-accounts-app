package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host     string
	Port     int
	DB       int
	Password string
}

// ConnectToRedis opens new connection with redis.
func ConnectToRedis(rc *RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rc.Host, rc.Port),
		Password: rc.Password,
		DB:       0,
	})

	return rdb
}
