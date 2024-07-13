package verifyCode

import (
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	aliyunUtil "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/luolayo/gin-study/global"
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

type aliyun struct {
	verificationCodeCache    *cache.Cache // 验证码 5 分钟过期
	verificationCodeReqCache *cache.Cache // 一分钟内只能发送一次验证码
}

var (
	aliyunOnce   sync.Once
	aliyunEntity *aliyun
)

func getAliyunEntity() *aliyun {
	aliyunOnce.Do(func() {
		aliyunEntity = new(aliyun)
		aliyunEntity.verificationCodeReqCache = cache.New(time.Minute, time.Minute)
		aliyunEntity.verificationCodeCache = cache.New(time.Minute*5, time.Minute*5)
	})
	return aliyunEntity
}

func (a *aliyun) SendVerificationCode(phoneNumber string) (err error) {
	// 验证是否可以获取验证码（1分钟有效期）
	_, found := a.verificationCodeReqCache.Get(phoneNumber)
	if found {
		err = errors.New("请勿重复发送验证码")
		return
	}

	// 生成验证码
	verifyCode := CreateRandCode()

	// 发送短信
	err = a.SendSms(a.getVerifyCodeReq(phoneNumber, verifyCode))
	if err != nil {
		return
	}

	// 验证码加入缓存
	a.verificationCodeReqCache.SetDefault(phoneNumber, 1)
	a.verificationCodeCache.SetDefault(phoneNumber, verifyCode)

	return
}

func (a *aliyun) CheckVerificationCode(phoneNumber, verificationCode string) (err error) {
	cacheCode, found := a.verificationCodeCache.Get(phoneNumber)
	if !found {
		err = errors.New("验证码已失效")
		return
	}

	cc, sure := cacheCode.(string)
	if !sure {
		err = errors.New("内部服务出错")
		return
	}
	if cc != verificationCode {
		err = errors.New("验证码输入错误")
		return
	}
	return
}

// CreateClient 可以上官网查看示例 https://next.api.aliyun.com/api/Dysmsapi/2017-05-25/SendSms?params={}
func (a *aliyun) CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func (a *aliyun) SendSms(req dysmsapi20170525.SendSmsRequest) (_err error) {
	client, _err := a.CreateClient(tea.String(global.Aliyun.AccessKeyID), tea.String(global.Aliyun.AccessKeySecret))
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

func (a *aliyun) getVerifyCodeReq(phoneNumber, code string) (req dysmsapi20170525.SendSmsRequest) {
	req = dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("阿里云短信测试"),
		TemplateCode:  tea.String("SMS_154950909"),
		PhoneNumbers:  tea.String(phoneNumber),
		TemplateParam: tea.String(`{"code":"` + code + `"}`),
	}
	return
}
