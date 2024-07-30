package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/api/install"
)

func UseInstallRoutes(r *gin.RouterGroup) {
	r.GET("/check", install.CheckStep)
	r.POST("/step1", install.Step1)
	r.POST("/step2", install.Step2)
	r.POST("/step3", install.Step3)
}
