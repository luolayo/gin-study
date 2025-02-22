package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/interceptor"
	"github.com/luolayo/gin-study/util/verifyCode"
)

// SentVerificationCode godoc
// @Summary SentVerificationCode
// @Description Sent verification code
// @Tags SMS
// @Schemes http https
// @Produce  json
// @Param phone_number query string true "Phone number" length(11) example(18888888888) Format(18888888888)
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @router /SMS/send [Get]
func SentVerificationCode(c *gin.Context) {
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
