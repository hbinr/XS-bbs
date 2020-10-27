package e

var MsgFlags = map[int]string{
	SUCCESS:        "操作成功",
	ERROR:          "操作失败",
	INVALID_PARAMS: "请求参数错误",

	ERROR_NOT_EXIST_USER: "该用户不存在",
	ERROR_EXIST_USER:     "该用户已存在",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
