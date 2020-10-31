package model

// CommunityDto 社区标签
type CommunityDto struct {
	CommunityID   int64  `json:"communityID"`   // 社区编号
	CommunityName string `json:"communityName"` // 社区名称
	Introduction  string `json:"introduction"`  // 社区介绍
}
