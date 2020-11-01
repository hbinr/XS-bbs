package e

import "errors"

var (
	ErrorUserNotLogin = errors.New(CodeNeedLogin.Msg())
	ErrUserExist      = errors.New(CodeUserExist.Msg())
	ErrUserNotExist   = errors.New(CodeUserNotExist.Msg())
	ErrEmailExist     = errors.New(CodeEmailExist.Msg())
	ErrInvalidID      = errors.New(CodeInvalidID.Msg())
)
