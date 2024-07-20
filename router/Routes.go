package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/api"
	"github.com/luolayo/gin-study/middleware"
)

func UserRoutes(r *gin.RouterGroup) {
	r.POST("/register", api.RegisterUser)
	r.GET("/checkPhone", api.CheckPhone)
	r.GET("/checkName", api.CheckName)
	r.GET("/info", middleware.Authentication(), api.UserInfo)
	r.POST("/login", api.UserLogin)
	r.GET("/logout", api.UserLogout)
}

func SMSRoutes(r *gin.RouterGroup) {
	r.GET("/send", api.SentVerificationCode)
}
