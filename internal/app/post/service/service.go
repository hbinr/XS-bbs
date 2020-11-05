package service

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	communityDao "xs.bbs/internal/app/community/dao"
	postDao "xs.bbs/internal/app/post/dao"
	"xs.bbs/internal/app/post/model"
	userDao "xs.bbs/internal/app/user/dao"
	"xs.bbs/internal/pkg/common"
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
		GetPostList(paging *common.PageInfo) ([]*model.PostDetailDto, int64, error)
		Vote(userID int64, voteP *model.PostVoteParam) error
	}
)

func NewPostService(db *gorm.DB, rdb *redis.Client) IPostService {
	return &postService{
		postDao:      postDao.NewPostDao(db, rdb),
		userDao:      userDao.NewUserDao(db),
		communityDao: communityDao.NewCommunityDao(db),
	}
}
