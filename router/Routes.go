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
	// Some APIs for obtaining content, the following APIs do not require login permission
	r.GET("/:id", api.GetContent)                          // get content by id
	r.GET("/post", api.GetPostContentList)                 // get post content list
	r.GET("/page", api.GetPageContentList)                 // get page content list
	r.GET("/attachment/:id", api.GetAttachmentContentList) // get attachment content list by post id or page id
	// Use middleware to check if the user is logged in.
	// The following API requires users to have login credentials
	r.Use(middleware.Authentication())
	r.PUT("/:id", api.UpdateContent)   // update content by id
	r.GET("/", api.GetUserContentList) // Query the articles logged in by the current user
	// Use middleware to check if the user is not a guest
	// The following API requires users to have normal user and above permissions
	r.Use(middleware.NotGustAuthority())
	r.POST("/", api.CreateContent)      // create content
	r.DELETE("/:id", api.DeleteContent) // delete content by id
	// use middleware to check if the user is an administrator
	// The following API requires users to have administrator privileges
	r.Use(middleware.AdminAuthority())
	r.GET("/approve/:id", api.ApproveRelease) // approve content by id

}
