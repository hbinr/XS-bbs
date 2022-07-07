package e

import "errors"

var (
	ErrUserNotLogin   = errors.New(CodeNeedLogin.Msg())
	ErrUserExist      = errors.New(CodeUserExist.Msg())
	ErrUserNotExist   = errors.New(CodeUserNotExist.Msg())
	ErrEmailExist     = errors.New(CodeEmailExist.Msg())
	ErrInvalidID      = errors.New(CodeInvalidID.Msg())
	ErrConvDataErr    = errors.New(CodeConvDataErr.Msg())
	ErrVoteTimeExpire = errors.New("投票时间已过")
)
