package middleware

import (
	"reflect"

	ut "github.com/go-playground/universal-translator"
	"xs.bbs/internal/app/user/model"
	"xs.bbs/internal/pkg/constant/key"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

//设置Translation
func Translation() gin.HandlerFunc {
	return func(c *gin.Context) {

		//设置支持语言
		en := en.New()
		zh := zh.New()

		//设置国际化翻译器
		uni := ut.New(zh, zh, en)
		val := validator.New()

		//根据参数取翻译器实例
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(locale)

		//翻译器注册到validator
		switch locale {
		case "en":
			enTranslations.RegisterDefaultTranslations(val, trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("en_comment")
			})
			break
		default:
			zhTranslations.RegisterDefaultTranslations(val, trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})

			model.UserParamValidAndTrans(val, trans)

			break
		}
		c.Set(key.TranslatorKey, trans)
		c.Set(key.ValidatorKey, val)
		c.Next()
	}
}
