package service

import (
	"github.com/google/wire"
	"xs.bbs/internal/app/community/dao"
	"xs.bbs/internal/app/community/model"
)

var _ ICommunityService = (*CommunityService)(nil)

// CommunityServiceSet CommunityServiceSet依赖注入
var CommunityServiceSet = wire.NewSet(
	wire.Struct(new(CommunityService), "*"),
	wire.Bind(new(ICommunityService), new(*CommunityService)),
)

type CommunityService struct {
	Dao dao.ICommunityDao
}

// ICommunityService 文章标签接口
type ICommunityService interface {
	// GetCommunityList 获取所有文章标签
	GetCommunityList() ([]model.CommunityDto, error)
	// GetCommunityDetailByID 根据社区id获取社区详情
	GetCommunityDetailByID(int64) (*model.CommunityDto, error)
}
