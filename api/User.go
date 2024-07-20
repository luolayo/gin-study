package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/interceptor"
	"github.com/luolayo/gin-study/model"
	"github.com/luolayo/gin-study/util"
	"github.com/luolayo/gin-study/util/verifyCode"
	"time"
)

// RegisterUser godoc
// @Summary RegisterUser
// @Description Register user
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param data body model.UserRegister true "User registration data"
// @Success 200 {object} interceptor.ResponseSuccess[model.User]
// @Failure 400 {object} interceptor.ResponseError
// @Failure 500 {object} interceptor.ResponseError
// @router /user/register [Post]
func RegisterUser(c *gin.Context) {

	userReigster := model.UserRegister{}
	if err := c.ShouldBind(&userReigster); err != nil {
		interceptor.BadRequest(c, "Invalid request", interceptor.ValidateErr(err))
		return
	}
	userModel := global.GormDB.Model(&model.User{})
	if userModel.Where("name = ?", userReigster.Name).First(&model.User{}).RowsAffected > 0 {
		interceptor.BadRequest(c, "User already exists", nil)
		return
	}
	if userModel.Where("phone = ?", userReigster.Phone).First(&model.User{}).RowsAffected > 0 {
		interceptor.BadRequest(c, "Phone number already exists", nil)
		return
	}
	if userReigster.Password != userReigster.ConfirmPassword {
		interceptor.BadRequest(c, "Password does not match", nil)
		return
	}
	if err := verifyCode.NewSms().CheckVerificationCode(userReigster.Phone, userReigster.Code); err != nil {
		interceptor.BadRequest(c, "Verification code error", nil)
		return
	}
	encryptPassword, err := util.Encrypt(userReigster.Password)
	if err != nil {
		interceptor.ServerError(c, "Failed to encrypt password")
		return
	}
	user := model.User{
		Name:       userReigster.Name,
		Password:   encryptPassword,
		Phone:      userReigster.Phone,
		Url:        userReigster.Url,
		ScreenName: userReigster.ScreenName,
	}
	tx := global.GormDB.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Failed to create user")
		return
	}
	if err := tx.Commit().Error; err != nil {
		interceptor.ServerError(c, "Failed to create user")
		return
	}
	token, _ := util.CreateToken(user)
	user.Token = token
	interceptor.Success(c, "Register success", user)
}

// CheckPhone godoc
// CheckPhone godoc
// @Summary CheckPhone
// @Description check phone number availability
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param phone query string true "UserPhone" length(11) example(18888888888) Format(18888888888)
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @router /user/checkPhone [Get]
func CheckPhone(c *gin.Context) {
	// Check phone number
	phone := c.Query("phone")
	if phone == "" {
		interceptor.BadRequest(c, "Phone number cannot be empty", nil)
		return
	}
	if global.GormDB.Model(&model.User{}).Where("phone = ?", phone).First(&model.User{}).RowsAffected > 0 {
		interceptor.BadRequest(c, "Phone number already exists", nil)
		return
	}
	interceptor.Success(c, "Phone number can be used", interceptor.Empty{})
}

// CheckName godoc
// RegisterUser godoc
// @Summary CheckName
// @Description check username availability
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param name query string true "Username" minlength(5)  maxlength(10) example(luolayo) Format(luolayo)
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @router /user/checkName [Get]
func CheckName(c *gin.Context) {
	// Check username
	name := c.Query("name")
	if name == "" {
		interceptor.BadRequest(c, "User name cannot be empty", nil)
		return
	}
	if global.GormDB.Model(&model.User{}).Where("name = ?", name).First(&model.User{}).RowsAffected > 0 {
		interceptor.BadRequest(c, "User name already exists", nil)
		return
	}
	interceptor.Success(c, "User name can be used", interceptor.Empty{})
}

// UserInfo godoc
// @Summary UserInfo
// @Description Get user information
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param Authorization header string true "Authorization token" example({{token}})
// @Success 200 {object} interceptor.ResponseSuccess[model.User]
// @Failure 401 {object} interceptor.ResponseError
// @Failure 400 {object} interceptor.ResponseError
// @router /user/info [Get]
func UserInfo(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		interceptor.Unauthorized(c, "Unauthorized")
		return
	}
	jwtClaims := claims.(util.JwtCustomClaims)
	user := model.User{}
	global.GormDB.Where("uid = ?", jwtClaims.ID).First(&user)
	if user.Uid == 0 {
		interceptor.Unauthorized(c, "Unauthorized")
		return
	}
	updateToken(&user, jwtClaims)
	global.LOG.Info("User information: %v", user)
	tx := global.GormDB.Begin()
	t := time.Now()
	user.Activated = &t
	global.LOG.Info("User information: %v", user)
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Failed to update user information")
		return
	}
	if err := tx.Commit().Error; err != nil {
		interceptor.ServerError(c, "Failed to update user information")
		return
	}
	interceptor.Success(c, "success", user)
}

// UserLogin godoc
// @Summary UserLogin
// @Description User login
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param data body model.UserLogin true "User login data"
// @Success 200 {object} interceptor.ResponseSuccess[model.User]
// @Failure 400 {object} interceptor.ResponseError
// @Failure 500 {object} interceptor.ResponseError
// @router /user/login [Post]
func UserLogin(c *gin.Context) {
	userLogin := model.UserLogin{}
	if err := c.ShouldBind(&userLogin); err != nil {
		interceptor.BadRequest(c, "Invalid request", interceptor.ValidateErr(err))
		return
	}
	user := model.User{}
	global.GormDB.Where("name = ?", userLogin.Name).First(&user)
	if user.Uid == 0 {
		interceptor.BadRequest(c, "User does not exist", nil)
		return
	}
	if ok, _ := util.Compare(user.Password, userLogin.Password); !ok {
		interceptor.BadRequest(c, "Password error", nil)
		return
	}
	token, err := util.CreateToken(user)
	if err != nil {
		global.LOG.Error("Failed to create token %v", err)
		interceptor.ServerError(c, "Failed to create token")
		return
	}
	user.Token = token
	// Update login time
	t := time.Now()
	user.Logged = &t
	global.GormDB.Save(&user)
	interceptor.Success(c, "Login success", user)
}

// UserLogout godoc
// @Summary UserLogout
// @Description User logout
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param Authorization header string true "Authorization token"
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @router /user/logout [Post]
func UserLogout(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		interceptor.Unauthorized(c, "Unauthorized")
		return
	}
	jwtClaims := claims.(util.JwtCustomClaims)
	user := model.User{}
	global.GormDB.Where("uid = ?", jwtClaims.ID).First(&user)
	if user.Uid == 0 {
		interceptor.Unauthorized(c, "Unauthorized")
		return
	}
	err := global.Redis.Del(jwtClaims.Name)
	if err != nil {
		global.LOG.Error("Failed to delete token %v", err)
		interceptor.ServerError(c, "Failed to delete token")
		return
	}
	interceptor.Success(c, "Logout success", interceptor.Empty{})
}

// updateToken Update token
func updateToken(user *model.User, token util.JwtCustomClaims) {
	newToken, err := util.UpdateToken(token)
	if err != nil {
		global.LOG.Error("Failed to update token %v", err)
		return
	}
	user.Token = newToken
}
