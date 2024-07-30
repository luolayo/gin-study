package verifyCode

import (
	"github.com/luolayo/gin-study/core"
	"github.com/luolayo/gin-study/global"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	core.InitViper("../../")
	global.InitRedis()
}

func TestAliyun_SendSms(t *testing.T) {
	sms := NewSms()
	err := sms.SendVerificationCode("188888888888")
	if err != nil {
		assert.Equal(t, "send sms failed", err.Error())
	}
}

func TestAliyun_CheckVerificationCode(t *testing.T) {
	sms := NewSms()
	err := sms.CheckVerificationCode("188888888888", "123456")
	if err != nil {
		assert.Equal(t, "验证码已失效", err.Error())
	}
}
