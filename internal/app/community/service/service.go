package service

import (
	"xs.bbs/internal/app/community/model"
	"xs.bbs/internal/app/community/repository"
)

var _ CommunityService = (*communityService)(nil)

// CommunityServiceSet CommunityServiceSet依赖注入
//var CommunityServiceSet = wire.NewSet(
//	wire.Struct(new(CommunityService), "*"),
//	wire.Bind(new(CommunityService), new(*CommunityService)),
//)

type (
	CommunityDto     = model.CommunityDto
	communityService struct {
		repo repository.CommunityRepo
	}

	// CommunityService 文章标签接口
	CommunityService interface {
		// GetCommunityList 获取所有文章标签
		GetCommunityList() ([]*CommunityDto, error)
		// GetCommunityDetailByID 根据社区id获取社区详情
		GetCommunityDetailByID(int64) (*CommunityDto, error)
	}
)

func NewCommunityService(repo repository.CommunityRepo) CommunityService {
	return &communityService{
		repo: repo,
	}
}
