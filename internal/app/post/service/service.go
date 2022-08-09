package service

import (
	"context"

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
		postRepo      post.PostRepo
		userRepo      user.UserRepo
		communityRepo community.CommunityRepo
	}
	PostService interface {
		Create(ctx context.Context, post *model.PostParam) error
		GetPostByID(ctx context.Context, id int64) (*model.PostDetailDto, error)
		GetPostListByIDs(ctx context.Context, paging *util.PageInfo) ([]*model.PostDetailDto, int64, error)
		Vote(ctx context.Context, userID int64, voteP *model.PostVoteParam) error
	}
)

func NewPostService(post post.PostRepo, user user.UserRepo,
	commu community.CommunityRepo) PostService {
	return &postService{
		postRepo:      post,
		userRepo:      user,
		communityRepo: commu,
	}
}
