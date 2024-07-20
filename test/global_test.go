package test

import (
	"github.com/luolayo/gin-study/core"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/util"
	"testing"
)

func TestPrintSystemInfo(t *testing.T) {
	core.InitGlobal()
	util.PrintSystemInfo()
	global.LOG.Info("test log")
	global.LOG.Error("test log")
	global.LOG.Warn("test log")
	global.LOG.Debug("test log")

}
