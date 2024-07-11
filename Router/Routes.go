package Router

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/Controller/Test"
)

func UserRoutes(r *gin.RouterGroup) {

}

func TestRoutes(r *gin.RouterGroup) {
	r.GET("", TestController.Ping)
	r.POST("", TestController.Pong)
}
