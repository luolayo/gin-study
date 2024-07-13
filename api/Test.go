package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/interceptor"
	"github.com/luolayo/gin-study/model"
	"github.com/luolayo/gin-study/util"
	"github.com/luolayo/gin-study/util/verifyCode"
)

// Ping godoc
// @Summary Ping
// @Description Test ping
// @Tags Test
// @Schemes http https
// @Produce  json
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @router /test [Get]
func Ping(c *gin.Context) {
	interceptor.Success(c, "success", gin.H{})
}

// Pong godoc
// @Summary Pong
// @Description Test pong
// @Tags Test
// @Schemes http https
// @Accept  json
// @Produce  json
// @Param data body model.Test true "Test data"
// @Success 200 {object} interceptor.ResponseSuccess[model.Test]
// @Failure 400 {object} interceptor.ResponseError
// @router /test [Post]
func Pong(c *gin.Context) {
	test := model.Test{}
	if err := c.ShouldBind(&test); err != nil {
		interceptor.BadRequest(c, "Invalid parameter", interceptor.ValidateErr(err))
		return
	}
	testModel := global.GormDB.Model(&test)
	testModel.Create(&test)
	interceptor.Success(c, "success", test)
}

// TestSentVerificationCode godoc
// @Summary TestSentVerificationCode
// @Description Sent verification code
// @Tags Test
// @Schemes http https
// @Produce  json
// @Param phone_number query string true "Phone number"
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @router /test/sentVerificationCode [Get]
func TestSentVerificationCode(c *gin.Context) {
	phoneNumber := c.Query("phone_number")
	if phoneNumber == "" {
		interceptor.BadRequest(c, "Invalid parameter", nil)
		return
	}
	err := verifyCode.NewSms().SendVerificationCode(phoneNumber)
	if err != nil {
		interceptor.BadRequest(c, "Failed to send verification code", nil)
		global.LOG.Error("Failed to send verification code %v", err)
		return
	}
	interceptor.Success(c, "success", gin.H{})
}

// TestCheckVerificationCode godoc
// @Summary TestCheckVerificationCode
// @Description Check verification code
// @Tags Test
// @Schemes http https
// @Produce  json
// @Param phone_number query string true "Phone number"
// @Param verification_code query string true "Verification code"
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @router /test/checkVerificationCode [Get]
func TestCheckVerificationCode(c *gin.Context) {
	phoneNumber := c.Query("phone_number")
	verificationCode := c.Query("verification_code")
	if phoneNumber == "" || verificationCode == "" {
		interceptor.BadRequest(c, "Invalid parameter", nil)
		return
	}

	if err := verifyCode.NewSms().CheckVerificationCode(phoneNumber, verificationCode); err != nil {
		interceptor.BadRequest(c, "Failed to check verification code", nil)
		return
	}
	interceptor.Success(c, "success", gin.H{})
}

// TestEncryption godoc
// @Summary TestEncryption
// @Description Test encryption function
// @Tags Test
// @Schemes http https
// @Accept  json
// @Produce  json
// @Param password query string true "Password"
// @Success 200 {object} interceptor.ResponseSuccess[string]
// @Failure 400 {object} interceptor.ResponseError
// @router /test/encryption [Get]
func TestEncryption(c *gin.Context) {
	password := c.Query("password")
	if password == "" {
		interceptor.BadRequest(c, "Invalid parameter", nil)
		return
	}
	encryption, err := util.Encrypt(password)
	if err != nil {
		interceptor.BadRequest(c, "Failed to encrypt", nil)
		global.LOG.Error("Failed to encrypt %v", err)
		return
	}
	interceptor.Success(c, "success", encryption)
}

// TestDecryption godoc
// @Summary TestDecryption
// @Description Test encryption function
// @Tags Test
// @Schemes http https
// @Accept  json
// @Produce  json
// @Param encryption query string true "encryption"
// @Success 200 {object} interceptor.ResponseSuccess[string]
// @Failure 400 {object} interceptor.ResponseError
// @router /test/decryption [Get]
func TestDecryption(c *gin.Context) {
	encryption := c.Query("encryption")
	if encryption == "" {
		interceptor.BadRequest(c, "Invalid parameter", nil)
		return
	}
	decryption, err := util.Decrypt(encryption)
	if err != nil {
		interceptor.BadRequest(c, "Failed to decrypt", nil)
		global.LOG.Error("Failed to decrypt %v", err)
		return
	}
	interceptor.Success(c, "success", decryption)
}
