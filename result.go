package rpclient

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/goccy/go-json"
	"gopkg.in/guregu/null.v4"
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

// ConvertDataTo 将 Data 数据提取到指定的结构体中
func (r Result) ConvertDataTo(dstPtr any) error {
	if dstPtr == nil || r.Data == nil {
		return nil
	}

	// 检查 dstPtr 是否为指针
	dstVo := reflect.ValueOf(dstPtr)
	if dstVo.Kind() != reflect.Ptr {
		return errors.New("rpclient: 'dstPtr' param type must be a pointer")
	}

	// 判断来源和目的数据类型是否可转换
	// map => struct|map
	// []map => []struct
	// 只在类型完全不兼容时才会报错
	srcKind := reflect.TypeOf(r.Data).Kind()
	dstKind := dstVo.Elem().Kind()
	if (srcKind == reflect.Map && dstKind != reflect.Struct && dstKind != reflect.Map) ||
		(srcKind == reflect.Slice && dstKind != reflect.Slice) {
		return fmt.Errorf("rpclient: %s cannot be converted to %s", srcKind, dstKind)
	}

	b, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, dstPtr)
}
