package TestController

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/Interceptor"
	"github.com/luolayo/gin-study/Model"
)

// Ping godoc
// @Summary Ping
// @Description Test ping
// @Tags Test
// @Schemes http https
// @Produce  json
// @Success 200 {object} Interceptor.ResponseSuccess[Interceptor.Empty]
// @Router /test [Get]
func Ping(c *gin.Context) {
	Interceptor.Success(c, "success", gin.H{})
}

// Pong godoc
// @Summary Pong
// @Description Test pong
// @Tags Test
// @Schemes http https
// @Accept  json
// @Produce  json
// @Param msg formData string true "msg" default(pong)
// @Success 200 {object} Interceptor.ResponseSuccess[Model.Test]
// @Failure 400 {object} Interceptor.ResponseError
// @Router /test [Post]
func Pong(c *gin.Context) {
	test := Model.Test{}
	if err := c.ShouldBind(&test); err != nil {
		Interceptor.BadRequest(c, "Invalid parameter", Interceptor.ValidateErr(err))
		return
	}
	Interceptor.Success(c, "success", test)
}
