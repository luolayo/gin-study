package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/api"
)

func UserRoutes(r *gin.RouterGroup) {

}

func TestRoutes(r *gin.RouterGroup) {
	r.GET("", api.Ping)
	r.POST("", api.Pong)
	r.GET("/sentVerificationCode", api.SentVerificationCode)
}
