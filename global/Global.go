package global

import (
	"github.com/luolayo/gin-study/config"
	"github.com/luolayo/gin-study/core"
	"gorm.io/gorm"
)

var (
	SysConfig *config.System
	LOG       *core.Logger
	GormDB    *gorm.DB
	Redis     *core.RedisClient
	Aliyun    *config.Aliyun
)

func Init() {
	SysConfig = config.GetSystemConfig()
	var level core.Level
	switch SysConfig.Environment {
	case "development":
		level = core.DebugLevel
	default:
		level = core.ErrorLevel
	}
	LOG = core.NewLogger(level)
	GormDB = core.GetGorm()
	Redis = core.NewRedisClient()
	Aliyun = config.GetAliYunConfig()
}
