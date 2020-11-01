package dao

import (
	"errors"

	"xs.bbs/internal/pkg/constant/e"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"xs.bbs/internal/app/community/model"
)

func (c *CommunityDao) GetCommunityList() (resList []model.Community, err error) {
	if err = c.DB.Find(&resList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("Dao.GetCommunityList no data", zap.Error(err))
			err = gorm.ErrRecordNotFound
		}
		zap.L().Error("Dao.GetCommunityList failed", zap.Error(err))
	}
	return resList, err
}

func (c *CommunityDao) GetCommunityDetailByID(ID int64) (res *model.Community, err error) {
	res = new(model.Community)
	if err = c.DB.Where("community_id", ID).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("Dao.GetCommunityList no data", zap.Error(err))
			err = e.ErrInvalidID
		}
		zap.L().Error("Dao.GetCommunityList failed", zap.Error(err))
	}
	return res, err
}
