package models

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()
var rdbClient *redis.Client
var redisEnable bool

type cacheDb struct{}

func (c cacheDb) Set(key string, value interface{}, expiration int) error {
	if redisEnable {
		v, err := json.Marshal(value)
		if err == nil {
			rdbClient.Set(ctx, key, string(v), time.Second*time.Duration(expiration))
			return nil
		}
		return err
	}
	return errors.New("redis功能未开启")
}

func (c cacheDb) Get(key string, obj interface{}) bool {
	if redisEnable {
		valueStr, err1 := rdbClient.Get(ctx, key).Result()
		if err1 == nil && valueStr != "" {
			err2 := json.Unmarshal([]byte(valueStr), obj)
			return err2 == nil
		}
		return false
	}
	return false
}

func (c cacheDb) Del(key string) bool {
	if redisEnable {
		_, err := rdbClient.Del(ctx, key).Result()
		return err == nil
	}
	return false
}

// 清除缓存
func (c cacheDb) FlushAll() {
	if redisEnable {
		rdbClient.FlushAll(ctx)
	}
}

var CacheDb = &cacheDb{}
