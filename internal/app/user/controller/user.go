package controller

import (
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
		ginx.FailWithMessage(errStr, c)
		return
	}

	if uDto, err = u.userService.SignUp(&uParam); err != nil {
		ginx.FailWithMessage(err.Error(), c)
		return
	}
	ginx.OkDetailed(uDto, "注册用户成功", c)
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
		ginx.FailWithMessage(e.GetMsg(e.INVALID_PARAMS), c)
		return
	}

	if uDto, err = u.userService.SelectById(int64(userID)); err != nil {
		ginx.FailWithMessage("获取用户失败", c)
		return
	}
	ginx.OkDetailed(&uDto, "获取用户成功", c)
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
		ginx.FailWithMessage(e.GetMsg(e.INVALID_PARAMS), c)
		return
	}

	if !u.userService.Delete(int64(userID)) {
		ginx.FailWithMessage("删除用户失败", c)
		return
	}
	ginx.OkWithMessage("删除用户成功", c)
}
