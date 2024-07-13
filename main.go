package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/config"
	"github.com/luolayo/gin-study/core"
	"github.com/luolayo/gin-study/docs"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/logger"
	"github.com/luolayo/gin-study/router"
	"github.com/luolayo/gin-study/util"
)

// @BasePath /
// @title Gin Study API
// @Version 1.0.1
// @Description Gin study is a small project for beginners to learn by writing a blog built by the gin framework.
// @Host localhost:8080
// @Schemes http https
func main() {
	InitGlobal()
	util.PrintSystemInfo()
	if global.SysConfig.Environment == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	docs.SwaggerInfo.Version = "1.0"
	r := router.GetRouter()
	err := r.Run(":" + global.SysConfig.Port)
	if err != nil {
		global.LOG.Error("Error starting server: %s", err)
	}

}
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
	global.GormDB = core.GetGorm()
	global.Aliyun = config.GetAliYunConfig()
}
