package service

import (
	"go.uber.org/zap"
	"xs.bbs/internal/app/user/dao"
	"xs.bbs/internal/app/user/model"
	"xs.bbs/pkg/tool/hash"
	"xs.bbs/pkg/tool/jwt"
	"xs.bbs/pkg/tool/snowflake"

	"github.com/gogf/gf/util/gconv"
)

// SignUp .
func (u *userService) SignUp(param *model.SignUpParam) (dto *UserDto, err error) {
	var uModel dao.UserModel
	if err = u.dao.CheckUserByUserName(param.Username); err != nil {
		return
	}
	if err = u.dao.CheckUserByEmail(param.Email); err != nil {
		return
	}
	if err = gconv.Struct(param, &uModel); err != nil {
		return
	}
	uModel.UserID = snowflake.GenID()
	// 密码加密
	uModel.Password = hash.MD5String(param.Password)
	if err = u.dao.Insert(&uModel); err != nil {
		return
	}

	if err = gconv.Struct(uModel, &dto); err != nil {
		return
	}
	return
}

// SignIn 登陆
func (u *userService) SignIn(signIn *model.SignInParam) (token string, err error) {
	var user *model.User
	// 获取用户信息
	if user, err = u.dao.GetUserByName(signIn.Username); err != nil {
		return
	}
	// 验证密码
	if user.Password != hash.MD5String(signIn.Password) {
		return
	}
	// 生成token
	return jwt.GenToken(user.UserID)
}

// Delete 根据用户ID删除用户
func (u *userService) Delete(userID int64) bool {
	return u.dao.Delete(userID)
}

// Update 根据用户ID修改用户
func (u *userService) Update(user *UserDto) error {
	var uModel model.User
	if err := gconv.Struct(user, &uModel); err != nil {
		return err
	}
	return u.dao.Update(&uModel)
}

// SelectByName 根据用户名查询用户
func (u *userService) SelectByName(userName string) (*UserDto, error) {
	uModel, err := u.dao.GetUserByName(userName)
	if err != nil {
		return nil, err
	}
	var uDto UserDto
	if err = gconv.Struct(uModel, &uDto); err != nil {
		return nil, err
	}
	return &uDto, nil
}

// SelectByID 根据用户ID查询用户
func (u *userService) SelectByID(userID int64) (*UserDto, error) {
	uModel, err := u.dao.GetUserByID(userID)
	if err != nil {
		zap.L().Error("userDao.GetUserByID", zap.Error(err), zap.Int64("userID", userID))
		return nil, err
	}
	var uDto UserDto
	if err = gconv.Struct(uModel, &uDto); err != nil {
		return nil, err
	}
	return &uDto, nil
}
