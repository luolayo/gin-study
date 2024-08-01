package user

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/interceptor/res"
	"github.com/luolayo/gin-study/interceptor/validate"
	"github.com/luolayo/gin-study/model"
	"github.com/luolayo/gin-study/util"
	"github.com/luolayo/gin-study/util/verifyCode"
	"github.com/spf13/cast"
	"time"
)

// CheckPhone godoc
// @Summary CheckPhone
// @Description Before registering a user, the front-end needs to check if the phone already exists. If the phone already exists, the front-end should prevent the use of that phone to continue registration in order to reduce API requests.
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param phone query string true "UserPhone" length(11) example(18888888888)
// @Success 200 {object} res.Response[res.Empty]
// @Failure 400 {object} res.ErrorRes[[]string]
// @router /user/checkPhone [Get]
func CheckPhone(c *gin.Context) {
	// Check phone number
	phone := c.Query("phone")
	if phone == "" {
		res.BadRequest(c, []string{"Phone number cannot be empty"})
		return
	}
	// Retrieve data from the database. If the data is not empty, it is registered
	if CheckPhoneService(phone) {
		res.BadRequest(c, []string{"Phone number already exists"})
		return
	}
	res.SuccessNoData(c)
}

// CheckName godoc
// @Summary CheckName
// @Description Before registering a user, the front-end needs to check if the username already exists. If the username already exists, the front-end should prevent the use of that username to continue registration in order to reduce API requests.
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param name query string true "Username" minlength(5)  maxlength(10) example(luolayo) Format(luolayo)
// @Success 200 {object} res.Response[res.Empty]
// @Failure 400 {object} res.ErrorRes[[]string]
// @router /user/checkName [Get]
func CheckName(c *gin.Context) {
	// Check username
	name := c.Query("name")
	if name == "" {
		res.BadRequest(c, []string{"Username cannot be empty"})
		return
	}
	// Retrieve data from the database. If the data is not empty, it is registered
	if CheckNameService(name) {
		res.BadRequest(c, []string{"Username already exists"})
		return
	}
	res.SuccessNoData(c)
}

// Register godoc
// @Summary Register
// @Description User Registration API. The user registration API is used to register a new user. be careful! The front-end should perform verification before requesting APIs, such as checking if the phone number and username already exist.
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param data body RegisterDTO true "User registration data"
// @Success 200 {object} res.Response[model.User]
// @Failure 400 {object} res.ErrorRes[[]string]
// @Failure 500 {object} res.ErrorRes[[]string]
// @router /user/register [Post]
func Register(c *gin.Context) {
	userReigster := RegisterDTO{}
	// Verify the request data
	if err := c.ShouldBind(&userReigster); err != nil {
		res.BadRequest(c, validate.Err(err))
		return
	}
	if CheckNameService(userReigster.Name) {
		res.BadRequest(c, []string{"Username already exists"})
		return
	}
	if CheckPhoneService(userReigster.Phone) {
		res.BadRequest(c, []string{"Phone number already exists"})
		return
	}
	// Check if the password is the same
	// These checks should be done by the front-end, with a secondary check by the back-end to reduce the possibility of errors
	if userReigster.Password != userReigster.ConfirmPassword {
		res.BadRequest(c, []string{"The password is inconsistent"})
		return
	}
	if err := verifyCode.NewSms().CheckVerificationCode(userReigster.Phone, userReigster.Code); err != nil {
		res.BadRequest(c, []string{"Verification code error"})
		return
	}
	// Encrypt the password.
	// The password encryption function is located in the util/jwt.go file, and the corresponding environment variables should be filled in according to this file
	encryptPassword, err := util.Encrypt(userReigster.Password)
	if err != nil {
		res.ServerError(c, "Failed to encrypt password")
		return
	}
	user := model.User{
		Name:       userReigster.Name,
		Password:   encryptPassword,
		Phone:      userReigster.Phone,
		Url:        userReigster.Url,
		ScreenName: userReigster.ScreenName,
		IP:         c.ClientIP(),
	}
	// Using transactions to handle database addition operations
	if err := CreateUserService(&user); err != nil {
		res.ServerError(c, "Failed to register")
	}
	// In theory, we should only create tokens when logging in, but to reduce API calls,
	// we also create tokens directly during registration and return them,
	// allowing users to use the system without logging in again
	token, _ := util.CreateToken(user)
	user.Token = token
	res.Success(c, user)
}

// Login godoc
// @Summary Login
// @Description User login
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param data body LoginDTO true "User login data"
// @Success 200 {object} res.Response[model.User]
// @Failure 400 {object} res.ErrorRes[[]string]
// @Failure 500 {object} res.ErrorRes[[]string]
// @router /user/login [Post]
func Login(c *gin.Context) {
	userLogin := LoginDTO{}
	// Verify the request data
	if err := c.ShouldBind(&userLogin); err != nil {
		res.BadRequest(c, validate.Err(err))
		return
	}
	// Check if the user exists
	user, err := GetUserServiceByName(userLogin.Name)
	if err != nil {
		res.BadRequest(c, []string{"User does not exist"})
		return
	}
	if ok, _ := util.Compare(user.Password, userLogin.Password); !ok {
		res.BadRequest(c, []string{"Password error"})
		return
	}
	token, err := util.CreateToken(user)
	if err != nil {
		global.LOG.Error("Failed to create token %v", err)
		res.ServerError(c, "Failed to create token")
		return
	}
	user.Token = token
	user.IP = c.ClientIP()
	// Update login time
	t := time.Now()
	user.Logged = &t
	err = UpdateUserService(&user)
	if err != nil {
		global.LOG.Error("Failed to update login time %v", err)
	}
	res.Success(c, user)
}

/**
 *  The following interface needs to be operated after logging in
 */

// Logout godoc
// @Summary Logout
// @Description User logout
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param Authorization header string true "Authorization token" example({{token}})
// @Success 200 {object} res.Response[res.Empty]
// @Failure 401 {object} res.ErrorRes[[]string]
// @Failure 400 {object} res.ErrorRes[[]string]
// @router /user/logout [Get]
func Logout(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		res.Unauthorized(c)
		return
	}
	jwtClaims := claims.(util.JwtCustomClaims)
	_, err := GetUserServiceByUid(uint(jwtClaims.ID))
	if err != nil {
		res.Unauthorized(c)
		return
	}
	// Delete token from redis
	err = global.Redis.Del(jwtClaims.Name)
	if err != nil {
		global.LOG.Error("Failed to delete token %v", err)
		res.ServerError(c, "Failed to delete token")
		return
	}
	res.SuccessNoData(c)
}

// GetInfo godoc
// @Summary GetInfo
// @Description Get user information
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param Authorization header string true "Authorization token" example({{token}})
// @Success 200 {object} res.Response[model.User]
// @Failure 401 {object} res.ErrorRes[[]string]
// @Failure 400 {object} res.ErrorRes[[]string]
// @router /user/info [Get]
func GetInfo(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		res.Unauthorized(c)
		return
	}
	jwtClaims := claims.(util.JwtCustomClaims)
	user, err := GetUserServiceByUid(uint(jwtClaims.ID))
	updateToken(&user, jwtClaims)
	if err != nil {
		res.Unauthorized(c)
		return
	}
	user.IP = c.ClientIP()
	t := time.Now()
	user.Logged = &t
	err = UpdateUserService(&user)
	if err != nil {
		global.LOG.Error("Failed to update login time %v", err)
	}
	res.Success(c, user)
}

// UpdateInfo godoc
// @Summary UpdateInfo
// @Description Update the user information API, which can partially transmit the information that needs to be updated, or transmit all the information that needs to be updated
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param data body UpdateDTO true "User update data"
// @Param Authorization header string true "Authorization token" example({{token}})
// @Success 200 {object} res.Response[model.User]
// @Failure 400 {object} res.ErrorRes[[]string]
// @Failure 401 {object} res.ErrorRes[[]string]
// @router /user/updateUserInfo [Put]
func UpdateInfo(c *gin.Context) {
	userUpdate := UpdateDTO{}
	if err := c.ShouldBind(&userUpdate); err != nil {
		res.BadRequest(c, validate.Err(err))
		return
	}
	claims, ok := c.Get("claims")
	if !ok {
		res.Unauthorized(c)
		return
	}
	jwtClaims := claims.(util.JwtCustomClaims)
	user, err := GetUserServiceByUid(uint(jwtClaims.ID))
	if err != nil {
		res.Unauthorized(c)
		return
	}
	// Update whatever content is uploaded
	if userUpdate.Url != "" {
		user.Url = userUpdate.Url
	}
	if userUpdate.ScreenName != "" {
		user.ScreenName = userUpdate.ScreenName
	}
	err = UpdateUserService(&user)
	if err != nil {
		res.ServerError(c, "Failed to update user information")
		return
	}
	res.Success(c, user)
}

// UpdatePassword godoc
// @Summary UpdatePassword
// @Description Before users can change their password, they need to send a verification code, which can only be updated after successful verification
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param Authorization header string true "Authorization token" example({{token}})
// @Param newPassword formData string true "New password" minlength(6)  maxlength(20) example(123456)
// @Param code formData string true "Verification code" length(6) example(123456)
// @Success 200 {object} res.Response[res.Empty]
// @Failure 400 {object} res.ErrorRes[[]string]
// @Failure 401 {object} res.ErrorRes[[]string]
// @router /user/updateUserPassword [Patch]
func UpdatePassword(c *gin.Context) {
	newPassword := c.PostForm("newPassword")
	// Updating password requires verifying phone number
	code := c.PostForm("code")
	if newPassword == "" {
		res.BadRequest(c, []string{"New password cannot be empty"})
		return
	}
	if code == "" {
		res.BadRequest(c, []string{"Verification code cannot be empty"})
		return
	}
	claims, ok := c.Get("claims")
	if !ok {
		res.Unauthorized(c)
		return
	}
	jwtClaims := claims.(util.JwtCustomClaims)
	user, err := GetUserServiceByUid(uint(jwtClaims.ID))
	if err != nil {
		res.Unauthorized(c)
		return
	}
	if err := verifyCode.NewSms().CheckVerificationCode(user.Phone, code); err != nil {
		res.BadRequest(c, []string{"Verification code error"})
		return
	}
	encryptPassword, err := util.Encrypt(newPassword)
	if err != nil {
		res.ServerError(c, "Failed to encrypt password")
		return
	}
	user.Password = encryptPassword
	err = UpdateUserService(&user)
	if err != nil {
		res.ServerError(c, "Failed to update password")
		return
	}
	res.SuccessNoData(c)
}

// UpdatePhone godoc
// @Summary UpdatePhone
// @Description Users need to verify the new phone number before updating their phone number
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param Authorization header string true "Authorization token" example({{token}})
// @Param phone formData string true "Phone number" length(11) example(18888888888) Format(18888888888)
// @Param code formData string true "Verification code" length(6) example(123456)
// @Success 200 {object} res.Response[res.Empty]
// @Failure 400 {object} res.ErrorRes[[]string]
// @Failure 401 {object} res.ErrorRes[[]string]
// @router /user/updateUserPhone [Patch]
func UpdatePhone(c *gin.Context) {
	phone := c.PostForm("phone")
	code := c.PostForm("code")
	if phone == "" {
		res.BadRequest(c, []string{"Phone number cannot be empty"})
		return
	}
	if code == "" {
		res.BadRequest(c, []string{"Verification code cannot be empty"})
		return
	}
	claims, ok := c.Get("claims")
	if !ok {
		res.Unauthorized(c)
		return
	}
	jwtClaims := claims.(util.JwtCustomClaims)
	user, err := GetUserServiceByUid(uint(jwtClaims.ID))
	if err != nil {
		res.Unauthorized(c)
		return
	}
	if err := verifyCode.NewSms().CheckVerificationCode(phone, code); err != nil {
		res.BadRequest(c, []string{"Verification code error"})
		return
	}
	user.Phone = phone
	err = UpdateUserService(&user)
	if err != nil {
		res.ServerError(c, "Failed to update phone number")
		return
	}
	res.SuccessNoData(c)
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
// @Success 200 {object} res.Response[model.User]
// @Failure 400 {object} res.ErrorRes[[]string]
// @Failure 401 {object} res.ErrorRes[[]string]
// @Failure 403 {object} res.ErrorRes[[]string]
// @router /user/getUserInfoById [Get]
func GetUserInfoById(c *gin.Context) {
	uid := c.Query("uid")
	if uid == "" {
		res.BadRequest(c, []string{"User id cannot be empty"})
		return
	}
	user, err := GetUserServiceByUid(cast.ToUint(uid))
	if err != nil {
		res.BadRequest(c, []string{"User does not exist"})
		return
	}
	res.Success(c, user)
}

// GetUserList godoc
// @Summary GetUserList
// @Description Due to the default registration of users as tourists, administrator review is required, and all users can be queried through this API
// @Tags User
// @Schemes http https
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param Authorization header string true "Authorization token" example({{token}})
// @Param pageSize query string false "Page size" example(10)
// @Param pageNum query string false "Page number" example(1)
// @Success 200 {object} res.Response[[]model.User]
// @Failure 400 {object} res.ErrorRes[[]string]
// @Failure 401 {object} res.ErrorRes[[]string]
// @Failure 403 {object} res.ErrorRes[[]string]
// @router /user/getUserList [Get]
func GetUserList(c *gin.Context) {
	pageSize := cast.ToInt(c.DefaultQuery("pageSize", "10"))
	pageNum := cast.ToInt(c.DefaultQuery("pageNum", "1"))
	users, err := GetUserServiceList(pageSize, pageNum)
	if err != nil {
		res.BadRequest(c, []string{"Failed to get user list"})
		return
	}
	res.Success(c, users)
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
// @Success 200 {object} res.Response[res.Empty]
// @Failure 400 {object} res.ErrorRes[[]string]
// @Failure 401 {object} res.ErrorRes[[]string]
// @Failure 403 {object} res.ErrorRes[[]string]
// @router /user/approveRegistration [Get]
func ApproveRegistration(c *gin.Context) {
	cid := c.Query("cid")
	if cid == "" {
		res.BadRequest(c, []string{"User id cannot be empty"})
		return
	}
	user, err := GetUserServiceByUid(cast.ToUint(cid))
	if err != nil {
		res.BadRequest(c, []string{"User does not exist"})
		return
	}
	user.Group = model.GroupUser
	err = UpdateUserService(&user)
	if err != nil {
		res.ServerError(c, "Failed to approve registration")
		return
	}
	res.SuccessNoData(c)
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
