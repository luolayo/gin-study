package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/api/SMS"
	"github.com/luolayo/gin-study/api/install"
	"github.com/luolayo/gin-study/api/test"
	"github.com/luolayo/gin-study/api/user"
	"github.com/luolayo/gin-study/middleware"
)

func UseInstallRoutes(r *gin.RouterGroup) {
	r.GET("/check", install.CheckStep)
	r.POST("/step1", install.Step1)
	r.POST("/step2", install.Step2)
	r.POST("/step3", install.Step3)
}

func UseUserRoutes(r *gin.RouterGroup) {
	// check routes
	r.GET("/checkPhone", user.CheckPhone) // check phoneNumber is existed
	r.GET("/checkName", user.CheckName)   // check username is existed
	// register and login router
	r.POST("/register", user.Register) // register user
	r.POST("/login", user.Login)       // login user
	// Use middleware to check if the user is logged in.
	r.Use(middleware.Authentication())
	// The following API requires users to have login credentials
	r.GET("/info", user.GetInfo)                        // get user info
	r.GET("/logout", user.Logout)                       // logout
	r.PUT("/update", user.UpdateInfo)                   // update user info
	r.PATCH("/updateUserPassword", user.UpdatePassword) // update user password
	r.PATCH("/updateUserPhone", user.UpdatePhone)       // update user phone
	// use middleware to check if the user is an administrator
	r.Use(middleware.AdminAuthority())
	// The following API requires users to have administrator privileges
	r.GET("/getUserInfoById", user.GetUserInfoById)         // Query user information through ID
	r.GET("/getUserList", user.GetUserList)                 // Query user list
	r.GET("/approveRegistration", user.ApproveRegistration) // Approve user registration and transition from tourist status to administrator status

}

func UseSMSRoutes(r *gin.RouterGroup) {
	r.GET("/send", SMS.SentVerificationCode)
}

func UseTestRoutes(r *gin.RouterGroup) {
	r.GET("/IP", test.GetIPAddredd)
}
