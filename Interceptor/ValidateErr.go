package Interceptor

import "github.com/go-playground/validator/v10"

type Validate struct {
	Key string
	Msg string
}

func formatValidateErr(errs error) []Validate {
	var validateErrs []Validate
	for _, err := range errs.(validator.ValidationErrors) {
		validateErrs = append(validateErrs, Validate{
			Key: err.Field(),
			Msg: err.Tag(),
		})
	}
	return validateErrs
}

func ValidateErr(err error) []string {
	errs := formatValidateErr(err)
	var msgArr []string
	for _, v := range errs {
		msgArr = append(msgArr, v.Key+formatMsg(v.Msg))
	}
	return msgArr
}

func formatMsg(msg string) string {
	switch msg {
	case "required":
		return "不能为空"
	case "min":
		return "太短"
	case "max":
		return "太长"
	default:
		return msg
	}
}
