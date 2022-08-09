package repository

import (
	"context"
	"errors"

	"xs.bbs/internal/app/community/model"
	"xs.bbs/internal/pkg/constant/e"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (c *communityRepo) GetCommunityList(ctx context.Context) (resList []model.Community, err error) {
	if err = c.db.WithContext(ctx).Find(&resList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("repo.GetCommunityList no data", zap.Error(err))
			err = gorm.ErrRecordNotFound
		}
		zap.L().Error("repo.GetCommunityList failed", zap.Error(err))
	}
	return
}

func (c *communityRepo) GetCommunityDetailByID(ctx context.Context, ID int64) (res *model.Community, err error) {
	res = new(model.Community)
	if err = c.db.WithContext(ctx).Where("community_id", ID).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("repo.GetCommunityList no data", zap.Error(err))
			err = e.ErrInvalidID
		}
		zap.L().Error("repo.GetCommunityList failed", zap.Error(err))
	}
	return
}
