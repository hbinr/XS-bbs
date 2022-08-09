package controller

import (
	"errors"

	"xs.bbs/internal/app/user/model"
	"xs.bbs/internal/pkg/constant/e"
	"xs.bbs/internal/pkg/ginx"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary 用户注册账号
// @Description 用户注册
// @Tags 用户接口
// @ID /user/signup
// @Accept  json
// @Produce  json
// @Param body body model.RegisterReq true "body"
// @Success 200 {object} ginx.Resp{data=model.UserDto} "success"
// @Router /user/signup [post]
func (u *UserController) Register(c *gin.Context) {
	var (
		err    error
		uParam model.RegisterReq
		uDto   *model.UserDto
	)
	if errStr := ginx.BindAndValid(c, &uParam); errStr != "" {
		ginx.RespErrorWithMsg(c, e.CodeError, errStr)
		return
	}

	uDto, err = u.userService.Register(c.Request.Context(), &uParam)

	switch err {
	case nil:
		ginx.RespSuccess(c, uDto)
	case e.ErrEmailExist:
		ginx.RespError(c, e.CodeEmailExist)
	case e.ErrConvDataErr:
		ginx.RespError(c, e.CodeConvDataErr)
	default:
		ginx.RespError(c, e.CodeError)
	}
}

// Login godoc
// @Summary 登录
// @Description 登录
// @Tags 用户接口
// @ID /user/signin
// @Accept  json
// @Produce json
// @Param body body model.LoginReq true "body参数"
// @Success 200 {string} string "ok" "登陆成功"
// @Router /user/signin [post]
func (u *UserController) Login(c *gin.Context) {
	var (
		err       error
		signParam model.LoginReq
		token     string
	)

	if errStr := ginx.BindAndValid(c, &signParam); errStr != "" {
		ginx.RespErrorWithMsg(c, e.CodeError, errStr)
		return
	}
	if token, err = u.userService.Login(c.Request.Context(), &signParam); err != nil {
		if errors.Is(err, e.ErrUserNotExist) {
			ginx.RespError(c, e.CodeUserNotExist)
			return
		}
		ginx.RespError(c, e.CodeWrongUserNameOrPassword)
		return
	}
	token, err = u.userService.Login(c.Request.Context(), &signParam)

	switch err {
	case nil:
		ginx.RespSuccess(c, token)
	case e.ErrUserNotExist:
		ginx.RespError(c, e.CodeUserNotExist)
	default:
		ginx.RespError(c, e.CodeWrongUserNameOrPassword)
	}
}

// Get godoc
// @Summary 根据id获取用户
// @Description 根据id获取用户
// @Tags 用户接口
// @ID /user/Get
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {object} ginx.Resp{data=model.UserDto} "success"
// @Router /user/Get [get]
func (u *UserController) Get(c *gin.Context) {
	var (
		userID int64
		err    error
		uDto   *model.UserDto
	)

	if userID, err = ginx.QueryInt("userID", c); err != nil {
		ginx.RespError(c, e.CodeInvalidParams)
		return
	}

	uDto, err = u.userService.SelectByID(c.Request.Context(), userID)

	switch err {
	case nil:
		ginx.RespSuccess(c, uDto)
	case e.ErrUserNotExist:
		ginx.RespError(c, e.CodeUserNotExist)
	case e.ErrConvDataErr:
		ginx.RespError(c, e.CodeConvDataErr)
	default:
		ginx.RespError(c, e.CodeError)
	}
}

// Delete godoc
// @Summary 根据id删除用户
// @Description 根据id删除用户
// @Tags 用户接口
// @ID /user/delete
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {object} ginx.Resp{data=string} "success"
// @Router /user/delete [get]
func (u *UserController) Delete(c *gin.Context) {
	var (
		userID int64
		err    error
	)

	if userID, err = ginx.QueryInt("userID", c); err != nil {
		ginx.RespError(c, e.CodeInvalidParams)
		return
	}

	if !u.userService.Delete(c.Request.Context(), userID) {
		ginx.RespError(c, e.CodeError)
		return
	}
	ginx.RespSuccess(c, nil)
}
