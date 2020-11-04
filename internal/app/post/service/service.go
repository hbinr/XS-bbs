package service

import (
	"gorm.io/gorm"
	communityDao "xs.bbs/internal/app/community/dao"
	postDao "xs.bbs/internal/app/post/dao"
	"xs.bbs/internal/app/post/model"
	userDao "xs.bbs/internal/app/user/dao"
)

var _ IPostService = (*postService)(nil)

//var PostServiceSet = wire.NewSet(
//	new(postService), "*",
//	wire.Bind(new(IPostService), new(*postService)),
//)

type (
	postService struct {
		postDao      postDao.IPostDao
		userDao      userDao.IUserDao
		communityDao communityDao.ICommunityDao
	}
	IPostService interface {
		Create(post *model.PostParam) error
		GetPostByID(pID int64) (*model.PostDetailDto, error)
	}
)

func NewPostService(db *gorm.DB) IPostService {
	return &postService{
		postDao:      postDao.NewPostDao(db),
		userDao:      userDao.NewUserDao(db),
		communityDao: communityDao.NewCommunityDao(db),
	}
}
