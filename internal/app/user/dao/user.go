package dao

import (
	"errors"

	"gorm.io/gorm"
	"xs.bbs/internal/app/user/model"
	"xs.bbs/internal/pkg/constant/e"
)

// Insert 新增用户
func (u *UserDao) Insert(user *model.User) (err error) {
	return u.db.Create(&user).Error
}

// Delete 根据用户ID删除用户，软删除
func (u *UserDao) Delete(userID int64) bool {
	return u.db.Where("user_id = ?", userID).Delete(&model.User{}).RowsAffected > 0
}

// Update 根据用户ID修改用户
func (u *UserDao) Update(user *model.User) error {
	return u.db.Where("user_id = ?").Updates(&user).Error
}

// GetUserByName 根据用户名查询用户
func (u *UserDao) GetUserByName(userName string) (user *model.User, err error) {
	user = new(model.User)
	if err = u.db.Where("username = ?", userName).Find(user).Error; err != nil {
		return
	}
	return
}

// GetUserByID 根据用户ID查询用户
func (u *UserDao) GetUserByID(userID int64) (user *model.User, err error) {
	user = new(model.User)
	if err = u.db.Where("user_id = ?", userID).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = e.ErrUserNotExist
			return
		}
		return
	}
	return
}

// CheckUserByUserName 根据userName检查用户是否存在
func (u *UserDao) CheckUserByUserName(userName string) error {
	var count int64
	if err := u.db.Model(&model.User{}).Where("username = ?", userName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return e.ErrUserExist
	}
	return nil
}

// CheckUserByEmail 通过email检查用户
func (u *UserDao) CheckUserByEmail(email string) error {
	var count int64
	if err := u.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return e.ErrEmailExist
	}
	return nil
}
