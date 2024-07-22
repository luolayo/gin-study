package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/api"
	"github.com/luolayo/gin-study/middleware"
)

func UserRoutes(r *gin.RouterGroup) {
	// check routes
	r.GET("/checkPhone", api.CheckPhone) // check phoneNumber is existed
	r.GET("/checkName", api.CheckName)   // check username is existed
	// register and login router
	r.POST("/register", api.RegisterUser) // register user
	r.POST("/login", api.UserLogin)       // login user
	// Use middleware to check if the user is logged in.
	r.Use(middleware.Authentication())
	// The following API requires users to have login credentials
	r.GET("/info", api.UserInfo)                           // get user info
	r.GET("/logout", api.UserLogout)                       // logout
	r.PUT("/update", api.UpdateUserInfo)                   // update user info
	r.PATCH("/updateUserPassword", api.UpdateUserPassword) // update user password
	r.PATCH("/updateUserPhone", api.UpdateUserPhone)       // update user phone
	// use middleware to check if the user is an administrator
	r.Use(middleware.AdminAuthority())
	// The following API requires users to have administrator privileges
	r.GET("/getUserInfoById", api.GetUserInfoById)         // Query user information through ID
	r.GET("/getUserList", api.GetUserList)                 // Query user list
	r.GET("/approveRegistration", api.ApproveRegistration) // Approve user registration and transition from tourist status to administrator status
}

func SMSRoutes(r *gin.RouterGroup) {
	r.GET("/send", api.SentVerificationCode)
}

func ContentRoutes(r *gin.RouterGroup) {
	r.GET("/:id", api.GetContent)
	r.GET("/post", api.GetPostContentList)
	r.GET("/page", api.GetPageContentList)
	r.GET("/attachment/:id", api.GetAttachmentContentList)
	r.Use(middleware.Authentication())
	r.PUT("/:id", api.UpdateContent)
	r.Use(middleware.NotGustAuthority())
	r.POST("/", api.CreateContent)
	r.DELETE("/:id", api.DeleteContent)
	r.Use(middleware.AdminAuthority())

}
