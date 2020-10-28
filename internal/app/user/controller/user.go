package controller

import (
	"errors"

	"xs.bbs/internal/app/user/dao"
	"xs.bbs/internal/app/user/model"
	"xs.bbs/internal/pkg/constant/e"
	"xs.bbs/internal/pkg/ginx"

	"github.com/gin-gonic/gin"
)

// SignUp godoc
// @Summary 用户注册账号
// @Description 用户注册
// @Tags 用户接口
// @ID /user/signup
// @Accept  json
// @Produce  json
// @Param body body model.SignUpParam true "body"
// @Success 200 {object} ginx.Response{data=model.UserDto} "success"
// @Router /user/signup [post]
func (u *UserController) SignUp(c *gin.Context) {
	var (
		err    error
		uParam model.SignUpParam
		uDto   *model.UserDto
	)
	if errStr := ginx.BindAndValid(c, &uParam); errStr != "" {
		ginx.ResponseErrorWithMsg(c, e.CODE_ERROR, errStr)
		return
	}

	if uDto, err = u.userService.SignUp(&uParam); err != nil {
		if errors.Is(err, dao.ErrUserExist) {
			ginx.ResponseError(c, e.CODE_USER_EXIST)
			return
		}
		if errors.Is(err, dao.ErrEmailExist) {
			ginx.ResponseError(c, e.CODE_EMAIL_EXIST)
			return
		}
		ginx.ResponseError(c, e.CODE_ERROR)
		return
	}
	ginx.ResponseSuccess(c, uDto)
}

// SignIn godoc
// @Summary 登录
// @Description 登录
// @Tags 用户接口
// @ID /user/signin
// @Accept  json
// @Produce json
// @Param body body model.SignInParam true "body参数"
// @Success 200 {string} string "ok" "登陆成功"
// @Router /user/signin [post]
func (u *UserController) SignIn(c *gin.Context) {
	var (
		err       error
		signParam model.SignInParam
	)
	if errStr := ginx.BindAndValid(c, &signParam); errStr != "" {
		ginx.ResponseErrorWithMsg(c, e.CODE_ERROR, errStr)
		return
	}

	if err = u.userService.SignIn(&signParam); err != nil {
		ginx.ResponseError(c, e.CODE_WRONG_USERNAME_OR_PASSWORD)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

// Get godoc
// @Summary 根据id获取用户
// @Description 根据id获取用户
// @Tags 用户接口
// @ID /user/Get
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {object} ginx.Response{data=model.UserDto} "success"
// @Router /user/Get [get]
func (u *UserController) Get(c *gin.Context) {
	var (
		userID int
		err    error
		uDto   *model.UserDto
	)

	if userID, err = ginx.QueryInt("userID", c); err != nil {
		ginx.ResponseError(c, e.CODE_INVALID_PARAMS)
		return
	}

	if uDto, err = u.userService.SelectByID(int64(userID)); err != nil {
		ginx.ResponseError(c, e.CODE_ERROR)
		return
	}
	ginx.ResponseSuccess(c, uDto)
}

// Delete godoc
// @Summary 根据id删除用户
// @Description 根据id删除用户
// @Tags 用户接口
// @ID /user/delete
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {object} ginx.Response{data=string} "success"
// @Router /user/delete [get]
func (u *UserController) Delete(c *gin.Context) {
	var (
		userID int
		err    error
	)

	if userID, err = ginx.QueryInt("userID", c); err != nil {
		ginx.ResponseError(c, e.CODE_INVALID_PARAMS)
		return
	}

	if !u.userService.Delete(int64(userID)) {
		ginx.ResponseError(c, e.CODE_ERROR)
		return
	}
	ginx.ResponseSuccess(c, nil)
}
