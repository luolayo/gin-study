package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/interceptor"
	"github.com/luolayo/gin-study/model"
	"github.com/luolayo/gin-study/util"
	"github.com/luolayo/gin-study/util/verifyCode"
)

// UserRegister godoc
// @Summary User registration
// @Description User registration
// @Tags User
// @Schemes http https
// @Accept  json
// @Produce  json
// @Param data body model.UserRegister true "User registration data"
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @router /user/register [Post]
func UserRegister(c *gin.Context) {
	userRegister := &model.UserRegister{}
	if err := c.ShouldBindJSON(userRegister); err != nil {
		interceptor.BadRequest(c, "Invalid parameter", interceptor.ValidateErr(err))
		return
	}
	// Check if the passwd and confirmPassword are consistent
	if userRegister.Passwd != userRegister.ConfirmPassword {
		interceptor.BadRequest(c, "Password inconsistency", nil)
		return
	}
	user := model.User{}
	global.GormDB.Where("phone = ?", userRegister.Phone).First(&user)
	if user.ID != 0 {
		interceptor.BadRequest(c, "User already exists", nil)
		return
	}
	// verification code
	if err := verifyCode.NewSms().CheckVerificationCode(userRegister.Phone, userRegister.Code); err != nil {
		interceptor.BadRequest(c, "Failed to check verification code", nil)
		return
	}
	// Generate ciphertext password
	passwd, _ := util.Encrypt(userRegister.Passwd)
	user = model.User{
		Name:   userRegister.Name,
		Phone:  userRegister.Phone,
		Passwd: passwd,
	}
	global.GormDB.Create(&user)
	interceptor.Success(c, "success", gin.H{})
}

// UserLogin godoc
// @Summary User login
// @Description User login
// @Tags User
// @Schemes http https
// @Accept  json
// @Produce  json
// @Param data body model.UserLogin true "User registration data"
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @router /user/login [Post]
func UserLogin(c *gin.Context) {
	userLogin := &model.UserLogin{}
	if err := c.ShouldBindJSON(userLogin); err != nil {
		interceptor.BadRequest(c, "Invalid parameter", interceptor.ValidateErr(err))
		return
	}
	user := model.User{}
	global.GormDB.Where("phone = ?", userLogin.Phone).First(&user)
	if user.ID == 0 {
		interceptor.BadRequest(c, "User does not exist", nil)
		return
	}
	if ok, _ := util.Compare(user.Passwd, userLogin.Passwd); !ok {
		interceptor.BadRequest(c, "Password error", nil)
		return
	}
	interceptor.Success(c, "success", gin.H{})
}
