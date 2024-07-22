package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.RecoveryMiddleware(global.LOG))
	r.Use(middleware.LoggerMiddleware(global.LOG))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router := r.Group("/user")
	UserRoutes(router)
	router = r.Group("/content")
	ContentRoutes(router)
	router = r.Group("/SMS")
	SMSRoutes(router)
	return r
}
