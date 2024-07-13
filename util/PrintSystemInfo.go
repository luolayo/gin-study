package util

import (
	"github.com/luolayo/gin-study/global"
	"os"
)

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
