package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetRouter() *gin.Engine {
	// GetRouter function, returns a gin.Engine instance
	r := gin.New()

	// swagger api docs route handler function ginSwagger.WrapHandler(swaggerFiles.Handler) is registered to the /swagger/*any route path using the GET method
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Use middleware.Cors() to enable cross-domain requests
	r.Use(middleware.Cors())

	// use middleware.RecoveryMiddleware() to recover from panics
	r.Use(middleware.RecoveryMiddleware(global.LOG))

	// use middleware.LoggerMiddleware() to log requests
	r.Use(middleware.LoggerMiddleware(global.LOG))

	// Use middleware.CheckInstall() to check if the application has been installed
	r.Use(middleware.CheckInstall())

	router := r.Group("install")
	UseInstallRoutes(router)

	router = r.Group("/SMS")
	UseSMSRoutes(router)

	router = r.Group("/user")
	UseUserRoutes(router)
	return r
}
