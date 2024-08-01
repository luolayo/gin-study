package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/core"
	_ "github.com/luolayo/gin-study/docs"
	"github.com/luolayo/gin-study/enum"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/router"
	"github.com/spf13/viper"
)

// @BasePath /
// @Title Gin Study API
// @Host localhost:8080
// @Version 1.1.0
// @Description Gin study is a small project for beginners to learn by writing a blog built by the gin framework.
// @Host localhost:8080
// @Schemes http https
func main() {
	// Check if the configuration file exists
	if !core.CheckConfigFile(enum.ConfigReleasePath) {
		// If the configuration file does not exist, create a new one and write the default configuration
		core.CreateConfigFile(enum.ConfigReleasePath)
		core.InitViper(enum.ConfigReleasePath)
		core.WriteBaseConfig()
	}
	// Initialize the configuration file
	core.InitViper(enum.ConfigReleasePath)
	// Determine whether the application is installed
	global.Init()
	if viper.GetBool("app.installed") {
		if err := global.AutoMigrate(); err != nil {
			panic(err)
		}
	}
	if viper.GetString("app.mode") == string(enum.Release) {
		gin.SetMode(gin.ReleaseMode)
	}
	r := router.GetRouter()
	err := r.Run(`:` + viper.GetString("app.port"))
	if err != nil {
		panic(err)
	}
}
