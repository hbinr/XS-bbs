package ginx

import (
	"net/http"

	"xs.bbs/internal/pkg/constant/e"

	"github.com/gin-gonic/gin"
)

// Response .
type Response struct {
	Code e.ResCode   `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// RespSuccess 响应成功
func RespSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: e.CodeSuccess,
		Msg:  e.CodeSuccess.Msg(),
		Data: data,
	})
}

// RespSuccess 响应失败，携带状态及对应信息
func RespError(c *gin.Context, code e.ResCode) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// RespSuccess 响应失败，携带状态+其他自定义信息
func RespErrorWithMsg(c *gin.Context, code e.ResCode, msg string) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
