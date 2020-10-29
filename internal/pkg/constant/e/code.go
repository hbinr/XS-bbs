package e

// ResCode .
type ResCode int

const (
	CodeSuccess                 ResCode = 200
	CodeInvalidParams           ResCode = 400
	CodeError                   ResCode = 500
	CodeConvDataErr             ResCode = 10000
	CodeValidateParamsErr       ResCode = 10001
	CodeUserNotExist            ResCode = 20001
	CodeUserExist               ResCode = 20002
	CodeEmailExist              ResCode = 20003
	CodeWrongPassword           ResCode = 20004
	CodeWrongUserNameOrPassword ResCode = 20005
)
