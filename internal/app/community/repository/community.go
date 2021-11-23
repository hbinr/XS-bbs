package repository

import (
	"errors"

	"xs.bbs/internal/app/community/model"
	"xs.bbs/internal/pkg/constant/e"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (c *communityRepo) GetCommunityList() (resList []model.Community, err error) {
	if err = c.db.Find(&resList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("dao.GetCommunityList no data", zap.Error(err))
			err = gorm.ErrRecordNotFound
		}
		zap.L().Error("dao.GetCommunityList failed", zap.Error(err))
	}
	return
}

func (c *communityRepo) GetCommunityDetailByID(ID int64) (res *model.Community, err error) {
	res = new(model.Community)
	if err = c.db.Where("community_id", ID).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("dao.GetCommunityList no data", zap.Error(err))
			err = e.ErrInvalidID
		}
		zap.L().Error("dao.GetCommunityList failed", zap.Error(err))
	}
	return
}
