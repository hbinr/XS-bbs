package e

var codeMsgMap = map[ResCode]string{
	CodeSuccess:       "操作成功",   // 200
	CodeInvalidParams: "请求参数错误", // 400
	CodeError:         "服务器繁忙",  // 500

	CodeConvDataErr:       "数据转换错误",   // 50000
	CodeValidateParamsErr: "参数校验错误",   // 50001
	CodeInvalidToken:      "无效的token", // 50002
	CodeNeedLogin:         "请先登陆",     // 50003
	CodeInvalidID:         "无效的ID",    // 50004

	CodeWrongPassword:           "密码错误",     // 40301
	CodeWrongUserNameOrPassword: "用户名或密码错误", // 40302
	CodeUserNotExist:            "该用户不存在",   // 40401
	CodeUserExist:               "该用户已存在",   // 40901
	CodeEmailExist:              "该邮箱已存在",   // 40902

}

// Msg .
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeError]
	}
	return msg
}
