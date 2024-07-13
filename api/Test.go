package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/interceptor"
	"github.com/luolayo/gin-study/model"
)

// Ping godoc
// @Summary Ping
// @Description Test ping
// @Tags Test
// @Schemes http https
// @Produce  json
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @router /test [Get]
func Ping(c *gin.Context) {
	interceptor.Success(c, "success", gin.H{})
}

// Pong godoc
// @Summary Pong
// @Description Test pong
// @Tags Test
// @Schemes http https
// @Accept  json
// @Produce  json
// @Param data body model.Test true "Test data"
// @Success 200 {object} interceptor.ResponseSuccess[model.Test]
// @Failure 400 {object} interceptor.ResponseError
// @router /test [Post]
func Pong(c *gin.Context) {
	test := model.Test{}
	if err := c.ShouldBind(&test); err != nil {
		interceptor.BadRequest(c, "Invalid parameter", interceptor.ValidateErr(err))
		return
	}
	interceptor.Success(c, "success", test)
}
