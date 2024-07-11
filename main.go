package main

import (
	"github.com/luolayo/gin-study/Core"
	"github.com/luolayo/gin-study/Logger"
	"github.com/luolayo/gin-study/Router"
)

func main() {
	system := Core.GetSystemConfig()
	router := Router.GetRouter()
	logger := Logger.NewLogger(Logger.InfoLevel)
	Core.PrintSystemInfo(logger, system)
	err := router.Run(":" + system.Port)
	if err != nil {
		logger.Error("Error starting server: %s", err)
	}

}
