package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/docs"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/router"
	"os"
)

// @BasePath /
// @title Gin Study API
// @Version 1.0.1
// @Description Gin study is a small project for beginners to learn by writing a blog built by the gin framework.
// @Host localhost:8080
// @Schemes http https
func main() {
	global.Init()
	PrintSystemInfo()
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

// PrintSystemInfo godoc
func PrintSystemInfo() {
	// Open banner file
	bannerFile, err := os.OpenFile("banner.txt", os.O_RDONLY, 0644)
	if err != nil {
		return
	}
	// read banner content
	banner := make([]byte, 1024)
	_, err = bannerFile.Read(banner)
	if err != nil {
		return
	}
	global.LOG.Info("\n" + string(banner))
	global.LOG.Info("%s Version %s", global.SysConfig.AppName, global.SysConfig.AppVersion)
	global.LOG.Info("Server running on http://127.0.0.1:%s", global.SysConfig.Port)
	defer func(bannerFile *os.File) {
		err := bannerFile.Close()
		if err != nil {
			return
		}
	}(bannerFile)
}
