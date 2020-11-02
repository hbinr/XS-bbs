package model

import "gorm.io/gorm"

// Post 帖子
type Post struct {
	gorm.Model
	PostID      int64  `gorm:"index:idx_post_id,unique"`      // 帖子id
	AuthorID    int64  `gorm:"index:idx_author_id,unique"`    // 作者的用户id
	CommunityID int64  `gorm:"index:idx_community_id,unique"` // 所属社区
	Status      int8   `gorm:"default:1"`                     // 帖子状态 1:有效,0:无效
	Title       string `gorm:"size:256"`                      // 标题
	Content     string `gorm:"type:text"`                     // 内容
}
