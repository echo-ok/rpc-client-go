package client

import (
	"errors"
	"fmt"
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
	if outputPtr == nil {
		return nil
	}

	// 检查 outputPtr 是否为指针
	outputPtrType := reflect.ValueOf(outputPtr)
	if outputPtrType.Kind() != reflect.Ptr {
		return errors.New("outputPtr 必须是一个指针")
	}

	sourceDataKind := reflect.TypeOf(r.Data).Kind()
	destinationDataKind := outputPtrType.Elem().Kind()
	if sourceDataKind != destinationDataKind {
		return fmt.Errorf("错误的类型：%s to %s", sourceDataKind.String(), destinationDataKind.String())
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
