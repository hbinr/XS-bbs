package dao

import (
	"time"

	"github.com/go-redis/redis"
	"xs.bbs/internal/pkg/constant/key"
	"xs.bbs/pkg/tool/snowflake"
)

func (p *postDao) Create(post *PostModel) (err error) {
	post.PostID = snowflake.GenID()
	// 1.存到MySQL中
	if err = p.db.Create(post).Error; err != nil {
		return
	}
	// 2.存到redis中
	pipeline := p.rdb.Pipeline()
	// 帖子时间
	pipeline.ZAdd(key.Redis(key.KeyPostTimeZset), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: post.PostID,
	})

	// 帖子分数
	pipeline.ZAdd(key.Redis(key.KeyPostScoreZset), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: post.PostID,
	})
	_, err = pipeline.Exec()
	return
}
func (p *postDao) GetPostByID(pID int64) (post *PostModel, err error) {
	post = new(PostModel)
	err = p.db.Where("post_id", pID).First(&post).Error
	return
}

func (p *postDao) GetPostList(limit, offset int) (posts []*PostModel, total int64, err error) {
	posts = make([]*PostModel, 0, limit) // 默认取limit条
	db := p.db.Model(&PostModel{})
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&posts).Error
	return
}
