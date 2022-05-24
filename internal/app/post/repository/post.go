package repository

import (
	"time"

	"xs.bbs/internal/pkg/constant"

	"github.com/go-redis/redis"
	"xs.bbs/internal/app/post/model"
	"xs.bbs/pkg/tool/snowflake"
)

func (p *postRepo) Create(post *model.Post) (err error) {
	post.PostID = snowflake.GenID()
	// 1.存到MySQL中
	if err = p.db.Create(post).Error; err != nil {
		return
	}
	// 2.存到redis中
	pipeline := p.rdb.Pipeline()
	// 帖子时间
	pipeline.ZAdd(constant.RedisKey(constant.KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: post.PostID,
	})

	// 帖子分数
	pipeline.ZAdd(constant.RedisKey(constant.KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: post.PostID,
	})
	_, err = pipeline.Exec()
	return
}

func (p *postRepo) GetPostByID(pID int64) (post *model.Post, err error) {
	post = new(model.Post)
	err = p.db.Where("post_id", pID).First(&post).Error
	return
}

func (p *postRepo) GetPostList(limit, offset int) (posts []*model.Post, total int64, err error) {
	posts = make([]*model.Post, 0, limit) // 默认取limit条
	db := p.db.Model(&model.Post{})
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&posts).Error
	return
}

// GetPostListByIDs 根据post_id切片获取post列表，并按照给定的post_id顺序返回
func (p *postRepo) GetPostListByIDs(pIDs []string) ([]*model.Post, int64, error) {
	panic("implement me")
}
