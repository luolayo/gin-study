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

// CheckPhone godoc
// @Summary CheckPhone
// @Description Before registering a user, the front-end needs to check if the phone already exists.
// @Description If the phone already exists, the front-end should prevent the use of that phone to continue registration in order to reduce API requests.
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param phone query string true "UserPhone" length(11) example(18888888888)
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
	// Retrieve data from the database. If the data is not empty, it is registered
	if global.GormDB.Model(&model.User{}).Where("phone = ?", phone).First(&model.User{}).RowsAffected > 0 {
		interceptor.BadRequest(c, "Phone number already exists", nil)
		return
	}
	interceptor.Success(c, "Phone number can be used", interceptor.Empty{})
}

// CheckName godoc
// @Summary CheckName
// @Description Before registering a user, the front-end needs to check if the username already exists.
// @Description If the username already exists, the front-end should prevent the use of that username to continue registration in order to reduce API requests.
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
	// Retrieve data from the database. If the data is not empty, it is registered
	if global.GormDB.Model(&model.User{}).Where("name = ?", name).First(&model.User{}).RowsAffected > 0 {
		interceptor.BadRequest(c, "User name already exists", nil)
		return
	}
	interceptor.Success(c, "User name can be used", interceptor.Empty{})
}

// RegisterUser godoc
// @Summary RegisterUser
// @Description User Registration API. The user registration API is used to register a new user.
// @Description be careful! The front-end should perform verification before requesting APIs, such as checking if the phone number and username already exist.
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
	// Verify the request data
	if err := c.ShouldBind(&userReigster); err != nil {
		interceptor.BadRequest(c, "Invalid request", interceptor.ValidateErr(err))
		return
	}
	// Check if user information is duplicated
	userModel := global.GormDB.Model(&model.User{})
	if userModel.Where("name = ?", userReigster.Name).First(&model.User{}).RowsAffected > 0 {
		interceptor.BadRequest(c, "User already exists", nil)
		return
	}
	if userModel.Where("phone = ?", userReigster.Phone).First(&model.User{}).RowsAffected > 0 {
		interceptor.BadRequest(c, "Phone number already exists", nil)
		return
	}
	// Check if the password is the same
	// These checks should be done by the front-end, with a secondary check by the back-end to reduce the possibility of errors
	if userReigster.Password != userReigster.ConfirmPassword {
		interceptor.BadRequest(c, "Password does not match", nil)
		return
	}
	if err := verifyCode.NewSms().CheckVerificationCode(userReigster.Phone, userReigster.Code); err != nil {
		interceptor.BadRequest(c, "Verification code error", nil)
		return
	}
	// Encrypt the password.
	// The password encryption function is located in the util/jwt.go file, and the corresponding environment variables should be filled in according to this file
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
	// Using transactions to handle database addition operations
	tx := global.GormDB.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Failed to create user")
		return
	}
	// If there are no issues, submit the transaction; if there are issues, roll back to ensure that the data is error free
	if err := tx.Commit().Error; err != nil {
		interceptor.ServerError(c, "Failed to create user")
		return
	}
	// In theory, we should only create tokens when logging in, but to reduce API calls,
	// we also create tokens directly during registration and return them,
	// allowing users to use the system without logging in again
	token, _ := util.CreateToken(user)
	user.Token = token
	interceptor.Success(c, "Register success", user)
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
	// Verify the request data
	if err := c.ShouldBind(&userLogin); err != nil {
		interceptor.BadRequest(c, "Invalid request", interceptor.ValidateErr(err))
		return
	}
	// Check if the user exists
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
	tx := global.GormDB.Begin()
	// Update the last active time
	t := time.Now()
	user.Activated = &t
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

// UserLogout godoc
// @Summary UserLogout
// @Description User logout
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param Authorization header string true "Authorization token" example({{token}})
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 401 {object} interceptor.ResponseError
// @Failure 400 {object} interceptor.ResponseError
// @router /user/logout [Get]
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
	// Delete token from redis
	err := global.Redis.Del(jwtClaims.Name)
	global.LOG.Info("Delete token %v", jwtClaims.Name)
	if err != nil {
		global.LOG.Error("Failed to delete token %v", err)
		interceptor.ServerError(c, "Failed to delete token")
		return
	}
	interceptor.Success(c, "Logout success", interceptor.Empty{})
}

// UpdateUserInfo godoc
// @Summary UpdateUserInfo
// @Description Update the user information API, which can partially transmit the information that needs to be updated, or transmit all the information that needs to be updated
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param data body model.UserUpdate true "User update data"
// @Param Authorization header string true "Authorization token" example({{token}})
// @Success 200 {object} interceptor.ResponseSuccess[model.User]
// @Failure 400 {object} interceptor.ResponseError
// @Failure 401 {object} interceptor.ResponseError
// @router /user/updateUserInfo [Put]
func UpdateUserInfo(c *gin.Context) {
	userUpdate := model.UserUpdate{}
	if err := c.ShouldBind(&userUpdate); err != nil {
		interceptor.BadRequest(c, "Invalid request", interceptor.ValidateErr(err))
		return
	}
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
	// Update whatever content is uploaded
	if userUpdate.Url != "" {
		user.Url = userUpdate.Url
	}
	if userUpdate.ScreenName != "" {
		user.ScreenName = userUpdate.ScreenName
	}
	tx := global.GormDB.Begin()
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

// UpdateUserPassword godoc
// @Summary UpdateUserPassword
// @Description Before users can change their password, they need to send a verification code, which can only be updated after successful verification
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param Authorization header string true "Authorization token" example({{token}})
// @Param newPassword formData string true "New password" minlength(6)  maxlength(20) example(123456)
// @Param code formData string true "Verification code" length(6) example(123456)
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @Failure 401 {object} interceptor.ResponseError
// @router /user/updateUserPassword [Patch]
func UpdateUserPassword(c *gin.Context) {
	newPassword := c.PostForm("newPassword")
	// Updating password requires verifying phone number
	code := c.PostForm("code")
	if newPassword == "" {
		interceptor.BadRequest(c, "New password cannot be empty", nil)
		return
	}
	if code == "" {
		interceptor.BadRequest(c, "Verification code cannot be empty", nil)
		return
	}
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
	if err := verifyCode.NewSms().CheckVerificationCode(user.Phone, code); err != nil {
		interceptor.BadRequest(c, "Verification code error", nil)
		return
	}
	encryptPassword, err := util.Encrypt(newPassword)
	if err != nil {
		interceptor.ServerError(c, "Failed to encrypt password")
		return
	}
	tx := global.GormDB.Begin()
	user.Password = encryptPassword
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Failed to update user password")
		return
	}
	if err := tx.Commit().Error; err != nil {
		interceptor.ServerError(c, "Failed to update user password")
		return
	}
	interceptor.Success(c, "success", interceptor.Empty{})
}

// UpdateUserPhone godoc
// @Summary UpdateUserPhone
// @Description Users need to verify the new phone number before updating their phone number
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param Authorization header string true "Authorization token" example({{token}})
// @Param phone formData string true "Phone number" length(11) example(18888888888) Format(18888888888)
// @Param code formData string true "Verification code" length(6) example(123456)
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @Failure 401 {object} interceptor.ResponseError
// @router /user/updateUserPhone [Patch]
func UpdateUserPhone(c *gin.Context) {
	phone := c.PostForm("phone")
	code := c.PostForm("code")
	if phone == "" {
		interceptor.BadRequest(c, "Phone number cannot be empty", nil)
		return
	}
	if code == "" {
		interceptor.BadRequest(c, "Verification code cannot be empty", nil)
		return
	}
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
	if err := verifyCode.NewSms().CheckVerificationCode(phone, code); err != nil {
		interceptor.BadRequest(c, "Verification code error", nil)
		return
	}
	tx := global.GormDB.Begin()
	user.Phone = phone
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Failed to update user phone number")
		return
	}
	if err := tx.Commit().Error; err != nil {
		interceptor.ServerError(c, "Failed to update user phone number")
		return
	}
	interceptor.Success(c, "success", interceptor.Empty{})
}

/*
 * The following APIs require administrator privileges and are operations performed by administrators
 */

// GetUserInfoById godoc
// @Summary GetUserInfoById
// @Description Get user information by id
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param uid query string true "User id" example(1)
// @Param Authorization header string true "Authorization token" example({{token}})
// @Success 200 {object} interceptor.ResponseSuccess[model.User]
// @Failure 400 {object} interceptor.ResponseError
// @Failure 401 {object} interceptor.ResponseError
// @Failure 403 {object} interceptor.ResponseError
// @router /user/getUserInfoById [Get]
func GetUserInfoById(c *gin.Context) {
	uid := c.Query("uid")
	if uid == "" {
		interceptor.BadRequest(c, "User id cannot be empty", nil)
		return
	}
	user := model.User{}
	global.GormDB.Where("uid = ?", uid).First(&user)
	if user.Uid == 0 {
		interceptor.BadRequest(c, "User does not exist", nil)
		return
	}
	interceptor.Success(c, "success", user)
}

// GetUserList godoc
// @Summary GetUserList
// @Description Due to the default registration of users as tourists, administrator review is required, and all users can be queried through this API
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param Authorization header string true "Authorization token" example({{token}})
// @Success 200 {object} interceptor.ResponseSuccess[[]model.User]
// @Failure 400 {object} interceptor.ResponseError
// @Failure 401 {object} interceptor.ResponseError
// @Failure 403 {object} interceptor.ResponseError
// @router /user/getUserList [Get]
func GetUserList(c *gin.Context) {
	var users []model.User
	// Query all users except administrators
	global.GormDB.Not(model.User{Group: model.GroupAdmin}).Find(&users)
	interceptor.Success(c, "success", users)
}

// ApproveRegistration godoc
// @Summary ApproveRegistration
// @Description Approve user registration
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param cid query string true "User id" example(1)
// @Param Authorization header string true "Authorization token" example({{token}})
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @Failure 401 {object} interceptor.ResponseError
// @Failure 403 {object} interceptor.ResponseError
// @router /user/approveRegistration [Get]
func ApproveRegistration(c *gin.Context) {
	cid := c.Query("cid")
	if cid == "" {
		interceptor.BadRequest(c, "User id cannot be empty", nil)
		return
	}
	user := model.User{}
	global.GormDB.Where("uid = ?", cid).First(&user)
	if user.Uid == 0 {
		interceptor.BadRequest(c, "User does not exist", nil)
		return
	}
	user.Group = model.GroupUser
	tx := global.GormDB.Begin()
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Failed to update user information")
		return
	}
	if err := tx.Commit().Error; err != nil {
		interceptor.ServerError(c, "Failed to update user information")
		return
	}
	interceptor.Success(c, "success", interceptor.Empty{})
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
