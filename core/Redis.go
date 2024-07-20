package core

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/luolayo/gin-study/config"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var (
	redisClient *RedisClient
	redisOnce   sync.Once
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient() *RedisClient {
	redisOnce.Do(func() {
		RedisConfig := config.GetRedisConfig()
		rdb := redis.NewClient(&redis.Options{
			Addr:         RedisConfig.Host + ":" + RedisConfig.Port,
			DB:           gconv.Int(RedisConfig.DB),
			DialTimeout:  time.Duration(gconv.Int(RedisConfig.DialTimeout)) * time.Second,
			ReadTimeout:  time.Duration(gconv.Int(RedisConfig.ReadTimeout)) * time.Second,
			WriteTimeout: time.Duration(gconv.Int(RedisConfig.WriteTimeout)) * time.Second,
			PoolSize:     gconv.Int(RedisConfig.PoolSize),
			PoolTimeout:  time.Duration(gconv.Int(RedisConfig.PoolTimeout)) * time.Second,
		})
		ctx := context.Background()
		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			panic(err)
		}

		redisClient = &RedisClient{
			client: rdb,
			ctx:    ctx,
		}
	})
	return redisClient
}

// Set 设置键值对
func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(r.ctx, key, value, expiration).Err()
}

// Get 获取键的值
func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

// Del 删除键
func (r *RedisClient) Del(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

// Close 关闭 Redis 客户端连接
func (r *RedisClient) Close() error {
	return r.client.Close()
}
