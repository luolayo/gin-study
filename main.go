package main

import (
	"github.com/luolayo/gin-study/Core"
	"github.com/luolayo/gin-study/Logger"
	"github.com/luolayo/gin-study/Router"
)

func main() {
	logger := Logger.NewLogger(Logger.InfoLevel)
	system := Core.GetSystemConfig()
	Core.PrintSystemInfo(logger, system)
	router := Router.GetRouter()
	err := router.Run(":" + system.Port)
	if err != nil {
		logger.Error("Error starting server: %s", err)
	}

}
