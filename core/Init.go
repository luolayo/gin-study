package core

import (
	"github.com/luolayo/gin-study/config"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/logger"
	"github.com/luolayo/gin-study/model"
	"gorm.io/gorm"
)

func InitGlobal() {
	global.SysConfig = config.GetSystemConfig()
	var level logger.Level
	switch global.SysConfig.Environment {
	case "development":
		level = logger.DebugLevel
	default:
		level = logger.ErrorLevel
	}
	global.LOG = logger.NewLogger(level)
	global.GormDB = GetGorm()
	global.Aliyun = config.GetAliYunConfig()
	if err := AutoMigrate(global.GormDB); err != nil {
		global.LOG.Error("AutoMigrate failed: %s", err)
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
		&model.Meta{},
		&model.Option{},
		&model.Relationship{},
	)
}
