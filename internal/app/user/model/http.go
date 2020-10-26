package model

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"xs.bbs/internal/pkg/ginx"
)

// SignUpParam 用户注册结构体
type SignUpParam struct {
	Username   string `json:"username" form:"username" comment:"用户名" validate:"required,valid_username"`
	Password   string `json:"password"`
	RePassword string `json:"rePassword"`
	NickName   string `json:"nickName"`
}

func (param *SignUpParam) BindValidParam(c *gin.Context) error {
	return ginx.DefaultGetValidParams(c, param)
}
func UserParamValidAndTrans(val *validator.Validate, trans ut.Translator) {
	//自定义验证方法
	val.RegisterValidation("valid_username", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "admin"
	})

	//自定义翻译器
	val.RegisterTranslation("valid_username", trans, func(ut ut.Translator) error {
		return ut.Add("valid_username", "{0} 填写不正确哦", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("valid_username", fe.Field())
		return t
	})
}
