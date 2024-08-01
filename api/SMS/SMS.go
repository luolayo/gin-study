package SMS

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/interceptor/res"
	"github.com/luolayo/gin-study/util/verifyCode"
)

// SentVerificationCode godoc
// @Summary SentVerificationCode
// @Description Sent verification code
// @Tags SMS
// @Schemes http https
// @Produce  json
// @Param phone_number query string true "Phone number" length(11) example(18888888888) Format(18888888888)
// @Success 200 {object} res.Response[model.User]
// @Failure 400 {object} res.ErrorRes[[]string]
// @router /SMS/send [Get]
func SentVerificationCode(c *gin.Context) {
	phoneNumber := c.Query("phone_number")
	if phoneNumber == "" {
		res.BadRequest(c, []string{"Invalid parameter"})
		return
	}
	err := verifyCode.NewSms().SendVerificationCode(phoneNumber)
	if err != nil {
		res.ServerError(c, "Failed to send verification code")
		global.LOG.Error("Failed to send verification code %v", err)
		return
	}
	res.SuccessNoData(c)
}
