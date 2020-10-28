package e

// ResCode .
type ResCode int

const (
	CODE_SUCCESS                    ResCode = 200
	CODE_INVALID_PARAMS             ResCode = 400
	CODE_ERROR                      ResCode = 500
	CODE_CONV_DATA_ERR              ResCode = 10000
	CODE_VALIDATE_PARAMS_ERR        ResCode = 10001
	CODE_USER_NOT_EXIST             ResCode = 20001
	CODE_USER_EXIST                 ResCode = 20002
	CODE_EMAIL_EXIST                ResCode = 20003
	CODE_WRONG_PASSWORD             ResCode = 20004
	CODE_WRONG_USERNAME_OR_PASSWORD ResCode = 20005
)
