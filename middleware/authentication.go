package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/interceptor"
	"github.com/luolayo/gin-study/model"
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
		v, _ := global.Redis.Get(claims.Name)
		if v != token {
			interceptor.Unauthorized(c, "Unauthorized")
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func AdminAuthority() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("claims")
		if !ok {
			interceptor.Unauthorized(c, "Unauthorized")
			c.Abort()
			return
		}
		jwtClaims := claims.(util.JwtCustomClaims)
		user := model.User{}
		global.GormDB.Where("uid = ?", jwtClaims.ID).First(&user)
		if user.Uid == 0 {
			interceptor.Unauthorized(c, "Unauthorized")
			c.Abort()
			return
		}
		if user.Group != model.GroupAdmin {
			interceptor.Forbidden(c, "Forbidden")
			c.Abort()
			return
		}
		c.Next()
	}
}

func NotGustAuthority() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("claims")
		if !ok {
			interceptor.Unauthorized(c, "Unauthorized")
			c.Abort()
			return
		}
		jwtClaims := claims.(util.JwtCustomClaims)
		user := model.User{}
		global.GormDB.Where("uid = ?", jwtClaims.ID).First(&user)
		if user.Uid == 0 {
			interceptor.Unauthorized(c, "Unauthorized")
			c.Abort()
			return
		}
		if user.Group == model.GroupGuest {
			interceptor.Forbidden(c, "Forbidden")
			c.Abort()
			return
		}
		c.Next()
	}
}
