package e

// ResCode .
type ResCode int

const (
	CODE_SUCCESS        ResCode = 200
	CODE_ERROR          ResCode = 500
	CODE_INVALID_PARAMS ResCode = 400

	ERROR_VALIDATE_PARAMS             ResCode = 10000
	ERROR_NOT_EXIST_USER              ResCode = 10001
	ERROR_EXIST_USER                  ResCode = 10002
	ERROR_EXIST_EMAIL                 ResCode = 10003
	ERROR_WRONG_USER_NAME_OR_PASSWORD ResCode = 10004
)
