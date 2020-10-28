package ginx

import (
	"net/http"

	"xs.bbs/internal/pkg/constant/e"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// Result  封装返回
func Result(code e.ResCode, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		int(code),
		data,
		msg,
	})
}

// Ok 响应成功
func Ok(c *gin.Context) {
	Result(e.CODE_SUCCESS, nil, "操作成功", c)
}

// OkWithMessage 响应成功且只携带信息
func OkWithMessage(message string, c *gin.Context) {
	Result(e.CODE_SUCCESS, nil, message, c)
}

// OkWithData 响应成功且只携带数据
func OkWithData(data interface{}, c *gin.Context) {
	Result(e.CODE_SUCCESS, data, e.CODE_SUCCESS.Msg(), c)
}

// OkDetailed 响应成功携带信息+数据
func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(e.CODE_SUCCESS, data, message, c)
}

// Fail 响应失败
func Fail(c *gin.Context) {
	Result(e.CODE_ERROR, nil, e.CODE_ERROR.Msg(), c)
}

// FailWithMessage 响应失败且携带错误信息
func FailWithMessage(message string, c *gin.Context) {
	Result(e.CODE_ERROR, nil, message, c)
}

// FailWithDetailed 响应失败且携带错误信息+数据
func FailWithDetailed(code e.ResCode, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}
