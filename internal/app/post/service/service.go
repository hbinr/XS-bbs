package service

import (
	"gorm.io/gorm"
	"xs.bbs/internal/app/post/dao"
	"xs.bbs/internal/app/post/model"
)

var _ IPostService = (*postService)(nil)

//var PostServiceSet = wire.NewSet(
//	new(postService), "*",
//	wire.Bind(new(IPostService), new(*postService)),
//)

type (
	PostDto     = model.PostDto
	postService struct {
		dao dao.IPostDao
	}
	IPostService interface {
		Create(post *model.PostParam) (*PostDto, error)
	}
)

func NewPostService(db *gorm.DB) IPostService {
	return &postService{
		dao: dao.NewPostDao(db),
	}
}
