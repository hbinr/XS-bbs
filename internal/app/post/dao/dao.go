package dao

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"xs.bbs/internal/app/post/model"
)

var _ IPostDao = (*postDao)(nil)

//var PostDaoSet = wire.NewSet(
//	new(postDao), "*",
//	wire.Bind(new(IPostDao), new(*postDao)),
//)

type (
	PostModel = model.Post
	postDao   struct {
		db  *gorm.DB
		rdb *redis.Client
	}

	IPostDao interface {
		Create(post *PostModel) error
		GetPostByID(pID int64) (*PostModel, error)
		GetPostList(page, pageSize int) ([]*PostModel, int64, error)
		Vote(userID, postID string, value float64) error
	}
)

func NewPostDao(db *gorm.DB, rdb *redis.Client) IPostDao {
	return &postDao{db: db, rdb: rdb}
}
