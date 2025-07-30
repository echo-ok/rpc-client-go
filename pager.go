package rpclient

// Pager 带翻页的数据列表
type Pager struct {
	Page       int  `json:"page"`         // 当前页码
	PageSize   int  `json:"page_size"`    // 每页数据条数
	TotalCount int  `json:"total_count"`  // 数据总数
	PageCount  int  `json:"page_count"`   // 总页数
	IsLastPage bool `json:"is_last_page"` // 是否最后一页
	Items      any  `json:"items"`        // 数据
}
