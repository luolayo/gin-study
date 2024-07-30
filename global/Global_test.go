package global

import (
	"github.com/luolayo/gin-study/core"
	"github.com/luolayo/gin-study/enum"
	"testing"
)

func init() {
	core.InitViper(enum.ConfigDevelopmentPath)
	Init()
}

func TestGlobal_DB(t *testing.T) {
	t.Log(DB.Name())
}

func TestGlobal_LOG(t *testing.T) {
	LOG.Info("test")
}

func TestGlobal_Redis(t *testing.T) {
	err := Redis.Set("test", "test", 0)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(Redis.Get("test"))
	defer func() {
		err := Redis.Close()
		if err != nil {
			t.Error(err)
			return
		}
	}()
}
