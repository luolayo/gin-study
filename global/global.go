package global

import (
	"github.com/luolayo/gin-study/config"
	"github.com/luolayo/gin-study/logger"
)

var (
	LOG       *logger.Logger
	SysConfig *config.System
)

func InitGlobal() {
	SysConfig = config.GetSystemConfig()
	var level logger.Level
	switch SysConfig.Environment {
	case "development":
		level = logger.DebugLevel
	default:
		level = logger.ErrorLevel
	}
	LOG = logger.NewLogger(level)
}
