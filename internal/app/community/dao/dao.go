package dao

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"xs.bbs/internal/app/community/model"
)

var _ ICommunityDao = (*CommunityDao)(nil)

// CommunityDaoSet CommunityDao依赖注入
var CommunityDaoSet = wire.NewSet(
	wire.Struct(new(CommunityDao), "*"),
	wire.Bind(new(ICommunityDao), new(*CommunityDao)))

type CommunityDao struct {
	DB *gorm.DB
}

// ICommunityDao 文章标签接口
type ICommunityDao interface {
	// GetCommunityList 获取所有文章标签
	GetCommunityList() ([]model.Community, error)
}
