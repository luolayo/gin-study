package core

import (
	"github.com/luolayo/gin-study/Config"
	"github.com/luolayo/gin-study/Logger"
	"os"
)

func PrintSystemInfo(logger *Logger.Logger, system Config.System) {
	// 读取banner
	bannerFile, err := os.OpenFile("banner.txt", os.O_RDONLY, 0644)
	if err != nil {
		return
	}
	banner := make([]byte, 1024)
	_, err = bannerFile.Read(banner)
	if err != nil {
		return
	}
	logger.Info("\n" + string(banner))
	logger.Info("%s Version %s", system.AppName, system.AppVersion)
	logger.Info("Server running on http://127.0.0.1:%s", system.Port)
	defer func(bannerFile *os.File) {
		err := bannerFile.Close()
		if err != nil {
			return
		}
	}(bannerFile)
}
