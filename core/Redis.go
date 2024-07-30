package core

import (
	"context"
	"github.com/luolayo/gin-study/config"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	client *redis.Client   // Redis client
	ctx    context.Context // Context
}

func NewRedisClient() *RedisClient {
	RedisConfig := config.GetRedisConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:         RedisConfig.Host + ":" + RedisConfig.Port,
		DB:           RedisConfig.DB,
		DialTimeout:  time.Duration(RedisConfig.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(RedisConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(RedisConfig.WriteTimeout) * time.Second,
		PoolSize:     RedisConfig.PoolSize,
		PoolTimeout:  time.Duration(RedisConfig.PoolTimeout) * time.Second,
	})
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return &RedisClient{
			client: nil,
			ctx:    ctx,
		}
	}

	return &RedisClient{
		client: rdb,
		ctx:    ctx,
	}
}

// Set key value pairs
func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(r.ctx, key, value, expiration).Err()
}

// Get the value of the key
func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

// Del delete key
func (r *RedisClient) Del(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

// Close the Redis client connection
func (r *RedisClient) Close() error {
	return r.client.Close()
}

func (r *RedisClient) CheckRedisConnection() bool {
	if r.client == nil {
		return false
	}
	return true
}
