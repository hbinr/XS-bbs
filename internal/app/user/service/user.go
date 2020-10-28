package service

import (
	"errors"

	"xs.bbs/internal/app/user/dao"
	"xs.bbs/internal/app/user/model"
	"xs.bbs/internal/pkg/constant/e"
	"xs.bbs/pkg/tool/hash"
	"xs.bbs/pkg/tool/snowflake"

	"github.com/gogf/gf/util/gconv"
)

// SignUp .
func (u *UserService) SignUp(param *model.SignUpParam) (dto *model.UserDto, err error) {
	var uModel dao.UserModel
	if err = u.Dao.CheckUserByUserName(param.Username); err != nil {
		return
	}
	if err = u.Dao.CheckUserByEmail(param.Email); err != nil {
		return
	}
	if err = gconv.Struct(param, &uModel); err != nil {
		return
	}
	uModel.UserID = snowflake.GenID()
	// 密码加密
	uModel.Password = hash.MD5String(param.Password)
	if err = u.Dao.Insert(&uModel); err != nil {
		return
	}

	if err = gconv.Struct(uModel, &dto); err != nil {
		return
	}
	return
}

// SignIn 登陆
func (u *UserService) SignIn(signIn *model.SignInParam) (err error) {
	var user *model.User

	if err = u.Dao.CheckUserByUserName(signIn.Username); err == nil {
		return errors.New(e.ERROR_NOT_EXIST_USER.Msg())
	}
	if user, err = u.Dao.SelectByName(signIn.Username); err != nil {
		return
	}
	if user.Password != hash.MD5String(signIn.Password) {
		return errors.New("密码错误")
	}
	return
}

// Delete 根据用户ID删除用户
func (u *UserService) Delete(userID int64) bool {
	return u.Dao.Delete(userID)
}

// Update 根据用户ID修改用户
func (u *UserService) Update(user *model.UserDto) error {
	var uModel model.User
	if err := gconv.Struct(user, &uModel); err != nil {
		return err
	}
	return u.Dao.Update(&uModel)
}

// SelectByName 根据用户名查询用户
func (u *UserService) SelectByName(userName string) (*model.UserDto, error) {
	uModel, err := u.Dao.SelectByName(userName)
	if err != nil {
		return nil, err
	}
	var uDto model.UserDto
	if err = gconv.Struct(uModel, &uDto); err != nil {
		return nil, err
	}
	return &uDto, nil
}

// SelectByID 根据用户ID查询用户
func (u *UserService) SelectByID(userID int64) (*model.UserDto, error) {
	uModel, err := u.Dao.SelectByID(userID)
	if err != nil {
		return nil, err
	}
	var uDto model.UserDto
	if err = gconv.Struct(uModel, &uDto); err != nil {
		return nil, err
	}
	return &uDto, nil
}
