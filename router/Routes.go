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
	r.POST("/login", api.UserLogin)
	r.GET("/info", middleware.Authentication(), api.UserInfo)
	r.GET("/logout", middleware.Authentication(), api.UserLogout)
	r.PUT("/update", middleware.Authentication(), api.UpdateUserInfo)
	r.PUT("/updateUserPassword", middleware.Authentication(), api.UpdateUserPassword)
	r.PUT("/updateUserPhone", middleware.Authentication(), api.UpdateUserPhone)
	r.GET("/getUserInfoById", middleware.Authentication(), middleware.AdminAuthority(), api.GetUserInfoById)
	r.GET("/getUserList", middleware.Authentication(), middleware.AdminAuthority(), api.GetUserList)
	r.GET("/approveRegistration", middleware.Authentication(), middleware.AdminAuthority(), api.ApproveRegistration)
}

func SMSRoutes(r *gin.RouterGroup) {
	r.GET("/send", api.SentVerificationCode)
}
