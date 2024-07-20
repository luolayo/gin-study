package test

import (
	"github.com/luolayo/gin-study/core"
	"github.com/luolayo/gin-study/util/verifyCode"
	"testing"
)

func TestSMS(t *testing.T) {
	core.InitGlobal()
	phone := "18888888888"
	err := verifyCode.NewSms().SendVerificationCode(phone)
	if err != nil {
		t.Error(err)
	}
	code := "123456"
	err = verifyCode.NewSms().CheckVerificationCode(phone, code)
}
