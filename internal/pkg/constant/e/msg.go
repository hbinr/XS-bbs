package e

var codeMsgMap = map[ResCode]string{
	CODE_SUCCESS:                    "操作成功",     // 200
	CODE_INVALID_PARAMS:             "请求参数错误",   // 400
	CODE_ERROR:                      "服务器繁忙",    // 500
	CODE_CONV_DATA_ERR:              "数据转换错误",   // 10000
	CODE_VALIDATE_PARAMS_ERR:        "参数校验错误",   // 10001
	CODE_USER_NOT_EXIST:             "该用户不存在",   // 20001
	CODE_USER_EXIST:                 "该用户已存在",   // 20002
	CODE_EMAIL_EXIST:                "该邮箱已存在",   // 20003
	CODE_WRONG_PASSWORD:             "密码错误",     // 20004
	CODE_WRONG_USERNAME_OR_PASSWORD: "用户名或密码错误", // 20005
}

// Msg .
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CODE_ERROR]
	}
	return msg
}
