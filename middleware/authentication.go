package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/api/user"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/interceptor/res"
	"github.com/luolayo/gin-study/model"
	"github.com/luolayo/gin-study/util"
	"github.com/spf13/cast"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			res.Unauthorized(c)
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			res.Unauthorized(c)
			c.Abort()
			return
		}
		v, _ := global.Redis.Get(claims.Name)
		if v != token {
			res.Unauthorized(c)
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
			res.Unauthorized(c)
			c.Abort()
			return
		}
		jwtClaims := claims.(util.JwtCustomClaims)
		userInfo, err := user.GetUserServiceByUid(cast.ToUint(jwtClaims.ID))
		if err != nil {
			res.Unauthorized(c)
			c.Abort()
			return
		}
		if userInfo.Group != model.GroupAdmin {
			res.Forbidden(c)
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
			res.Unauthorized(c)
			c.Abort()
			return
		}
		jwtClaims := claims.(util.JwtCustomClaims)
		userInfo, err := user.GetUserServiceByUid(cast.ToUint(jwtClaims.ID))
		if err != nil {
			res.Unauthorized(c)
			c.Abort()
			return
		}
		if userInfo.Group == model.GroupGuest {
			res.Forbidden(c)
			c.Abort()
			return
		}
		c.Next()
	}
}
