package cache

import (
	"github.com/go-redis/redis"
	"FreightDistribution/logger"
	"os"
	"strconv"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("CACHE_REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("CACHE_REDIS_ADDR"),
		Password: os.Getenv("CACHE_REDIS_PW"),
		DB:       int(db),
	})

	_, err := client.Ping().Result()

	if err != nil {
		logger.Log().Panic("连接Redis不成功", err)
	}

	RedisClient = client
}
