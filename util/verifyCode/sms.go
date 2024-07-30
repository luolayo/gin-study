package verifyCode

import (
	"fmt"
	"math/rand"
	"time"
)

type SmsOperation interface {
	SendVerificationCode(phoneNumber string) error
	CheckVerificationCode(phoneNumber, verificationCode string) error
}

func NewSms() SmsOperation {
	return getAliyunEntity()
}

// CreateRandCode create a random number
func CreateRandCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
