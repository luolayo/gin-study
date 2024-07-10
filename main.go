package main

import (
	"github.com/luolayo/gin-study/Logger"
	"github.com/luolayo/gin-study/Router"
)

func main() {
	router := Router.GetRouter()
	logger := Logger.NewLogger(Logger.InfoLevel)
	logger.Info("Server started on port 8080")
	err := router.Run(":8080")
	if err != nil {
		logger.Error("Error starting server: %s", err)
	}

}
