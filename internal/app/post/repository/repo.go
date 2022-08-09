package repository

import (
	"context"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
	"xs.bbs/internal/app/post/model"
)

var _ PostRepo = (*postRepo)(nil)

//var PostDaoSet = wire.NewSet(
//	new(postRepo), "*",
//	wire.Bind(new(PostRepo), new(*postRepo)),
//)

type (
	postRepo struct {
		db  *gorm.DB
		rdb *redis.Client
	}

	PostRepo interface {
		Create(ctx context.Context, post *model.Post) error
		GetPostByID(ctx context.Context, pID int64) (*model.Post, error)
		GetPostListByIDs(ctx context.Context, ids []string) ([]*model.Post, int64, error)
		GetPostList(ctx context.Context, page, pageSize int) ([]*model.Post, int64, error)
		// Vote 投票,数据存储于redis中
		Vote(ctx context.Context, userID, postID string, value float64) error
	}
)

func NewPostRepo(db *gorm.DB, rdb *redis.Client) PostRepo {
	return &postRepo{db: db, rdb: rdb}
}
