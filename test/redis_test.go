package test

import (
	"errors"
	"github.com/luolayo/gin-study/global"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	global.Init()
	err := global.Redis.Set("key", "value", time.Hour)
	if err != nil {
		t.Error("Error setting key:", err)
		return
	}

	value, err := global.Redis.Get("key")
	if (err != nil) && (!errors.Is(err, redis.Nil)) {
		t.Error("Error setting key:", err)
		return
	}
	t.Log("key:", value)

	// 删除键
	err = global.Redis.Del("key")
	if err != nil {
		t.Error("Error setting key:", err)
		return
	}
}
