package repository

import (
	"context"

	"gorm.io/gorm"
	"xs.bbs/internal/app/community/model"
)

var _ CommunityRepo = (*communityRepo)(nil)

// CommunityDaoSet CommunityDao依赖注入
//var CommunityDaoSet = wire.NewSet(
//	wire.Struct(new(communityRepo), "*"),
//	wire.Bind(new(CommunityRepo), new(*communityRepo)))

type (
	communityRepo struct {
		db *gorm.DB
	}

	// CommunityRepo 文章标签接口
	CommunityRepo interface {
		GetCommunityList(context.Context) ([]model.Community, error)
		GetCommunityDetailByID(context.Context, int64) (*model.Community, error)
	}
)

func NewCommunityRepo(db *gorm.DB) CommunityRepo {
	return &communityRepo{db: db}
}
