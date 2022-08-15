package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var RedisCache = &redis.Client{}

func init() {
	fmt.Println("init redis")
	RedisCache = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	result, err := RedisCache.Ping().Result()
	if err != nil {
		fmt.Println("ping error", err.Error())
		return
	}
	fmt.Println("ping result:", result)
}

// Get 获取缓存数据
func Get(key string) (string, error) {
	result, err := RedisCache.Get(key).Result()
	return result, err
}

// Set 设置数据 过期时间默认24H
func Set(key, value string) error {
	err := RedisCache.Set(key, value, time.Hour*24).Err()
	return err
}
func LPush(key string, values ...interface{}) error {
	err := RedisCache.LPush(key, values...).Err()
	return err
}
func RPop(key string) (string, error) {
	result, err := RedisCache.RPop(key).Result()
	return result, err
}
