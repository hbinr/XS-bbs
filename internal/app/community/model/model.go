package model

import "gorm.io/gorm"

// Community 社区标签
type Community struct {
	gorm.Model
	CommunityID   int64  `gorm:"not null;index:idx_community_id;"` // 社区编号
	CommunityName string `gorm:"size:32"`                          // 社区名称
	Introduction  string `gorm:"type:text"`                        // 社区介绍
}
