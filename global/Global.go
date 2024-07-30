package global

import (
	"github.com/luolayo/gin-study/config"
	"github.com/luolayo/gin-study/core"
	"github.com/luolayo/gin-study/enum"
	"github.com/luolayo/gin-study/model"
)

var (
	LOG   *core.Logger
	DB    *core.GormClient
	Redis *core.RedisClient
)

func Init() {
	InitLog()
	InitDB()
	InitRedis()
}

func InitDB() {
	DB = core.GetGorm()
}

func InitRedis() {
	Redis = core.NewRedisClient()
}

func InitLog() {
	appconfig := config.GetAppConfig()
	if appconfig.Mode == enum.Release {
		LOG = core.NewLogger(core.ErrorLevel)
	} else {
		LOG = core.NewLogger(core.DebugLevel)
	}
}

func AutoMigrate() error {
	client := DB.GetClient()
	return client.AutoMigrate(
		&model.Comment{},
		&model.Content{},
		&model.Link{},
		&model.Meta{},
		&model.Option{},
		&model.Relationship{},
		&model.User{},
	)
}
