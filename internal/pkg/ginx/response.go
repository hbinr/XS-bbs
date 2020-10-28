package ginx

import (
	"net/http"

	"xs.bbs/internal/pkg/constant/e"

	"github.com/gin-gonic/gin"
)

// Response .
type Response struct {
	Code e.ResCode   `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// ResponseSuccess 响应成功
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: e.CODE_SUCCESS,
		Msg:  e.CODE_SUCCESS.Msg(),
		Data: data,
	})
}

// ResponseError 响应失败，携带状态及对应信息
func ResponseError(c *gin.Context, code e.ResCode) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// ResponseErrorWithMsg 响应失败，携带状态+其他自定义信息
func ResponseErrorWithMsg(c *gin.Context, code e.ResCode, msg string) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
