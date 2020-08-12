package response

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var (
	trans ut.Translator
)

func init() {
	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})
		err := zh_translations.RegisterDefaultTranslations(validate, trans)
		if err != nil {
			fmt.Printf("Register default translations fail err: %s", err.Error())
		}
	}
}

func Translate(err validator.ValidationErrors) string {
	msg := err[0].Translate(trans)
	return msg
}
