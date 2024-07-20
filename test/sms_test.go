package test

import (
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/util/verifyCode"
	"testing"
)

func TestSMS(t *testing.T) {
	global.Init()
	phone := "18888888888"
	err := verifyCode.NewSms().SendVerificationCode(phone)
	if err != nil {
		t.Error(err)
	}
	code := "123456"
	err = verifyCode.NewSms().CheckVerificationCode(phone, code)
}
