package e

var codeMsgMap = map[ResCode]string{
	CodeSuccess:       "操作成功",   // 200
	CodeInvalidParams: "请求参数错误", // 400
	CodeError:         "服务器繁忙",  // 500

	CodeConvDataErr:       "数据转换错误",   // 10000
	CodeValidateParamsErr: "参数校验错误",   // 10001
	CodeInvalidToken:      "无效的token", // 10002
	CodeNeedLogin:         "请先登陆",     // 10003

	CodeUserNotExist:            "该用户不存在",   // 20001
	CodeUserExist:               "该用户已存在",   // 20002
	CodeEmailExist:              "该邮箱已存在",   // 20003
	CodeWrongPassword:           "密码错误",     // 20004
	CodeWrongUserNameOrPassword: "用户名或密码错误", // 20005
}

// Msg .
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeError]
	}
	return msg
}
