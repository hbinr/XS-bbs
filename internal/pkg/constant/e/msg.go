package e

var codeMsgMap = map[ResCode]string{
	CODE_SUCCESS:        "操作成功",
	CODE_ERROR:          "服务器繁忙",
	CODE_INVALID_PARAMS: "请求参数错误",

	ERROR_VALIDATE_PARAMS:             "参数校验错误",
	ERROR_NOT_EXIST_USER:              "该用户不存在",
	ERROR_EXIST_USER:                  "该用户已存在",
	ERROR_EXIST_EMAIL:                 "该邮箱已存在",
	ERROR_WRONG_USER_NAME_OR_PASSWORD: "用户名或密码错误",
}

// Msg .
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CODE_ERROR]
	}
	return msg
}
