package api

import (
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"regexp"
)

type MobileValidator struct {
}

//注册手机验证
func (m MobileValidator) Register(v *validator.Validate, trans ut.Translator) (err error) {
	//注册验证器
	err = v.RegisterValidation("mobile", mobile)

	//注册错误返回
	err = v.RegisterTranslation("mobile", trans, func(ut ut.Translator) error {
		return ut.Add("mobile", "手机号码格式不正确", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mobile", fe.Field())
		return t
	})
	return
}

//手机号码验证
func mobile(fl validator.FieldLevel) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)

	if fl.Field().String() == "" {
		return true
	} else {
		return rgx.MatchString(fl.Field().String())
	}
}
