package cache

import (
	"fmt"
	"time"

	"github.com/erfanwd/golang-course-project/config"
	"github.com/go-redis/redis/v7"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password:           cfg.Redis.Password,
		DB:                 0,
		DialTimeout:        cfg.Redis.DialTimeout * time.Second,
		ReadTimeout:        cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout:       cfg.Redis.WriteTimeout * time.Second,
		PoolSize:           cfg.Redis.PoolSize,
		PoolTimeout:        cfg.Redis.PoolTimeout * time.Second,
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: cfg.Redis.IdleCheckFrequency * time.Millisecond,
	})
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis(){
	redisClient.Close()
}
