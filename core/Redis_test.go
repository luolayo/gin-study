package core

import (
	"github.com/luolayo/gin-study/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

var redisC *RedisClient

// Some initialization done before testing
func init() {
	InitViper(enum.ConfigDevelopmentPath)
	redisC = NewRedisClient()
}

// TestRedisClient tests the RedisClient method
// Test whether the return of client is empty when the connection fails due to configuration errors or lack of configuration
func TestRedisClient_ConfigErr(t *testing.T) {
	if redisC.client != nil {
		t.Error("redisC.client is not nil")
	}
}

// TestRedisClient_Set tests the Set method
func TestRedisClient_Set(t *testing.T) {
	if redisC.client == nil {
		t.Error("redisC.client is nil")
		return
	}
	if err := redisC.Set("test", "test", 0); err != nil {
		t.Error(err)
	}
	defer func() {
		err := redisC.Close()
		if err != nil {
			t.Error(err)
		}
	}()
}

// TestRedisClient_Get tests the Get method
func TestRedisClient_Get(t *testing.T) {
	value, err := redisC.Get("test")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "test", value)
	defer func() {
		err := redisC.Close()
		if err != nil {
			t.Error(err)
		}
	}()
}

// TestRedisClient_Del tests the Del method
func TestRedisClient_Del(t *testing.T) {
	err := redisC.Del("test")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		err := redisC.Close()
		if err != nil {
			t.Error(err)
		}
	}()
}
