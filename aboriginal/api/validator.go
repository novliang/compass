package api

import (
	"errors"
	cn "github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"strings"
)

func init() {
	trans := setTranslator()
	registerCustomerValidator(trans)
}

type MyValidator struct {
	validator *validator.Validate
}

type ValidatorError struct {
	errors []string
}

func (err ValidatorError) Error() string {
	return strings.Join(err.errors, " ")
}

var Validator = MyValidator{
	validator: validator.New(),
}

//注册自定义验证器
func registerCustomerValidator(trans ut.Translator) {
	err := MobileValidator{}.Register(Validator.validator, trans)
	if err != nil {
		panic(err)
	}
}

//设置验证翻译语言
func setTranslator() ut.Translator {
	locale := cn.New()
	uni := ut.New(locale, locale)
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	err := zh_translations.RegisterDefaultTranslations(Validator.validator, trans)
	if err != nil {
		panic(err)
	}
	return trans
}

//定义验证方法
func (v *MyValidator) Validate(i interface{}) (err error) {
	trans := setTranslator()
	err = v.validator.Struct(i)
	if err != nil {
		ef := err.(validator.ValidationErrors)
		errs := ef.Translate(trans)
		r := []string{}
		for _, v := range errs {
			r = append(r, v)
		}
		err = errors.New(strings.Join(r, ","))
	}
	return
}
