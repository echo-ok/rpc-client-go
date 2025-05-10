package client

import (
	"errors"
	"github.com/goccy/go-json"
	"gopkg.in/guregu/null.v4"
	"reflect"
)

type Result struct {
	StoreId   string      `json:"store_id"`
	StoreName string      `json:"store_name"`
	Key       string      `json:"key"`
	Label     null.String `json:"label"`
	Ok        bool        `json:"ok"`
	Error     null.String `json:"error"`
	Data      any         `json:"data"`
}

// ConvertDataTo 将 Data 数据提取到结构体
func (r Result) ConvertDataTo(outputPtr any) error {
	// 检查 outputPtr 是否为指针
	if reflect.ValueOf(outputPtr).Kind() != reflect.Ptr {
		return errors.New("outputPtr 必须是一个指针")
	}

	b, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, outputPtr)
}

type ResultPagerData struct {
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
	TotalCount int  `json:"total_count"`
	PageCount  int  `json:"page_count"`
	IsLastPage bool `json:"is_last_page"`
	Items      any  `json:"items"`
}
