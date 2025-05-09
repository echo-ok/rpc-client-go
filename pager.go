package client

// Pager 带翻页的数据
type Pager struct {
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
	TotalCount int  `json:"total_count"`
	PageCount  int  `json:"page_count"`
	IsLastPage bool `json:"is_last_page"`
	Items      any  `json:"items"`
}

func NewPager(page, pageSize, total int) *Pager {
	if pageSize < 1 {
		pageSize = 10
	}

	pageCount := -1
	if total >= 0 {
		pageCount = (total + pageSize - 1) / pageSize
		if page > pageCount {
			page = pageCount
		}
	}
	if page < 1 {
		page = 1
	}

	if pageCount < 0 {
		pageCount = 0
	}

	return &Pager{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: total,
		PageCount:  pageCount,
		IsLastPage: page >= pageCount,
	}
}

func (p *Pager) SetItems(items any) *Pager {
	p.Items = items
	return p
}
