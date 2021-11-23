package common

// PageInfo 分页请求数据
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页条数
	Order    string `json:"order" form:"order"`       // 排序方式
}

func (p *PageInfo) Offset() int {
	offset := 0
	if p.Page > 0 {
		offset = (p.Page - 1) * p.PageSize
	}
	return offset
}

// PageResult 分页返回数据
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
