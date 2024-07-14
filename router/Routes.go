package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/api"
	"github.com/luolayo/gin-study/middleware"
)

func UserRoutes(r *gin.RouterGroup) {
	r.POST("/register", api.UserRegister)
	r.POST("/login", api.UserLogin)
	r.GET("/info", middleware.Authentication(), api.UserInfo)

}

func TestRoutes(r *gin.RouterGroup) {
	r.GET("", api.Ping)
	r.POST("", api.Pong)
	r.GET("/sentVerificationCode", api.TestSentVerificationCode)
	r.GET("/checkVerificationCode", api.TestCheckVerificationCode)
	r.GET("/encryption", api.TestEncryption)
	r.GET("/decryption", api.TestDecryption)
}

func UtilRoutes(r *gin.RouterGroup) {
	r.GET("/sentVerificationCode", api.SentVerificationCode)
}
