package Router

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	r := gin.Default()
	router := r.Group("/test")
	TestRoutes(router)
	router = r.Group("/user")
	UserRoutes(router)
	return r
}
