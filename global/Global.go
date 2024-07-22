package global

import (
	"github.com/luolayo/gin-study/config"
	"github.com/luolayo/gin-study/core"
	"github.com/luolayo/gin-study/model"
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
	if err := AutoMigrate(GormDB); err != nil {
		LOG.Error("AutoMigrate failed: %s", err)
		panic(err)
	}
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Test{},
		&model.User{},
		&model.Content{},
		&model.Comment{},
		&model.Link{},
		&model.Option{},
	)
}
