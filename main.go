package main

import (
	"github.com/luolayo/gin-study/Logger"
	"github.com/luolayo/gin-study/Router"
	"github.com/luolayo/gin-study/core"
)

func main() {
	system := core.GetSystemConfig()
	router := Router.GetRouter()
	logger := Logger.NewLogger(Logger.InfoLevel)
	core.PrintSystemInfo(logger, system)
	err := router.Run(":" + system.Port)
	if err != nil {
		logger.Error("Error starting server: %s", err)
	}

}
