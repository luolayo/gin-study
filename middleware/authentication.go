package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/interceptor"
	"github.com/luolayo/gin-study/util"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			interceptor.Unauthorized(c, "Unauthorized")
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			interceptor.Unauthorized(c, "Unauthorized")
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
