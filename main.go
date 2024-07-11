package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/Core"
	"github.com/luolayo/gin-study/Logger"
	"github.com/luolayo/gin-study/Router"
	"github.com/luolayo/gin-study/docs"
)

// @BasePath /
// @title Gin Study API
// @Version 1.0.1
// @Description Gin study is a small project for beginners to learn by writing a blog built by the gin framework.
// @Host localhost:8080
// @Schemes http https
func main() {
	logger := Logger.NewLogger(Logger.InfoLevel)
	system := Core.GetSystemConfig()
	Core.PrintSystemInfo(logger, system)
	if system.Environment == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	docs.SwaggerInfo.Version = "1.0"
	router := Router.GetRouter()
	err := router.Run(":" + system.Port)
	if err != nil {
		logger.Error("Error starting server: %s", err)
	}

}
