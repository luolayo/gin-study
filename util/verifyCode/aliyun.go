package verifyCode

import (
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	aliyunUtil "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/luolayo/gin-study/global"
	"github.com/spf13/viper"
	"sync"
	"time"
)

type aliyun struct {
}

var (
	aliyunOnce   sync.Once
	aliyunEntity *aliyun
)

// getAliyunEntity initializes and returns a singleton instance of aliyun
func getAliyunEntity() *aliyun {
	aliyunOnce.Do(func() {
		aliyunEntity = new(aliyun)
	})
	return aliyunEntity
}

// SendVerificationCode sends a verification code to the specified phone number
func (a *aliyun) SendVerificationCode(phoneNumber string) (err error) {
	// Check if a verification code was sent recently (within 1 minute)
	ok, _ := global.Redis.Get(phoneNumber)
	if ok != "" {
		err = errors.New("please do not send duplicate verification codes")
		return
	}

	// Generate a new verification code

	verifyCode := "123456"
	// Send the SMS
	if viper.GetString("app.mode") != "development" {
		verifyCode = CreateRandCode()
		err = a.SendSms(a.getVerifyCodeReq(phoneNumber, verifyCode))
	}
	if err != nil {
		return
	}

	// Store the verification code in the caches
	err = global.Redis.Set(phoneNumber, verifyCode, time.Minute)
	if err != nil {
		return err
	}
	return
}

// CheckVerificationCode checks if the provided verification code matches the stored code
func (a *aliyun) CheckVerificationCode(phoneNumber, verificationCode string) (err error) {
	cacheCode, found := global.Redis.Get(phoneNumber)
	if found != nil {
		err = errors.New("验证码已失效")
		return
	}

	cc := cacheCode
	if cc != verificationCode {
		err = errors.New("验证码输入错误")
		return
	}
	// Verification code is correct, remove it from the cache
	err = global.Redis.Del(phoneNumber)
	if err != nil {
		return err
	}
	return
}

// CreateClient creates a new Aliyun SMS client with the given access keys
func (a *aliyun) CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// Your AccessKey ID
		AccessKeyId: accessKeyId,
		// Your AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Set the endpoint for the Aliyun SMS service
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

// SendSms sends an SMS using the Aliyun SMS service
func (a *aliyun) SendSms(req dysmsapi20170525.SendSmsRequest) (_err error) {
	if viper.GetString("aliyun.accessKeyId") == "" || viper.GetString("aliyun.accessKeySecret") == "" {
		return errors.New("missing required parameters")
	}
	client, _err := a.CreateClient(tea.String(viper.GetString("aliyun.accessKeyId")), tea.String(viper.GetString("aliyun.accessKeySecret")))
	if _err != nil {
		return _err
	}

	defer func() {
		if r := tea.Recover(recover()); r != nil {
			_err = r
		}
	}()

	runtime := &aliyunUtil.RuntimeOptions{}
	result, _err := client.SendSmsWithOptions(&req, runtime)
	if _err != nil {
		return _err
	}

	if *result.Body.Code != "OK" {
		_err = errors.New(result.String())
		return
	}

	return _err
}

// getVerifyCodeReq creates a SendSmsRequest with the given phone number and verification code
func (a *aliyun) getVerifyCodeReq(phoneNumber, code string) (req dysmsapi20170525.SendSmsRequest) {
	req = dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(viper.GetString("aliyun.signName")),     // The SMS signature
		TemplateCode:  tea.String(viper.GetString("aliyun.templateCode")), // The SMS template code
		PhoneNumbers:  tea.String(phoneNumber),                            // The recipient's phone number
		TemplateParam: tea.String(`{"code":"` + code + `"}`),              // The verification code to be sent
	}
	return
}
