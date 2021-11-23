package service

import (
	community "xs.bbs/internal/app/community/repository"
	"xs.bbs/internal/app/post/model"
	post "xs.bbs/internal/app/post/repository"
	user "xs.bbs/internal/app/user/repository"
	"xs.bbs/internal/pkg/util"
)

var _ PostService = (*postService)(nil)

//var PostServiceSet = wire.NewSet(
//	new(postService), "*",
//	wire.Bind(new(PostService), new(*postService)),
//)

type (
	postService struct {
		postDao      post.PostRepo
		userDao      user.UserRepo
		communityDao community.CommunityRepo
	}
	PostService interface {
		Create(post *model.PostParam) error
		GetPostByID(pID int64) (*model.PostDetailDto, error)
		GetPostListByIDs(paging *util.PageInfo) ([]*model.PostDetailDto, int64, error)
		Vote(userID int64, voteP *model.PostVoteParam) error
	}
)

func NewPostService(post post.PostRepo, user user.UserRepo,
	commu community.CommunityRepo) PostService {
	return &postService{
		postDao:      post,
		userDao:      user,
		communityDao: commu,
	}
}
