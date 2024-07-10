package TestController

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/Logger"
	"github.com/luolayo/gin-study/Model"
)

func Ping(c *gin.Context) {
	c.String(200, "pong")
}

func Pong(c *gin.Context) {
	test := Model.Test{}
	logger := Logger.NewLogger(Logger.ErrorLevel)
	if err := c.ShouldBind(&test); err != nil {
		logger.Error("Error: %s", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": &test.Msg})
}
