package global

import (
	"github.com/luolayo/gin-study/config"
	"github.com/luolayo/gin-study/logger"
	"gorm.io/gorm"
)

var (
	LOG       *logger.Logger
	SysConfig *config.System
	GormDB    *gorm.DB
)
