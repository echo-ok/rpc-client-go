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
// 因为使用的是 json.Unmarshal 所以请确保 `json` 标签的正确性，否则可能会导致数据丢失
func (r Result) ConvertDataTo(dstPtr any) error {
	if dstPtr == nil {
		return errors.New("rpclient: 'dstPtr' param value cannot be nil")
	}

	if r.Data == nil {
		return nil
	}

	// 检查 dstPtr 是否为指针
	dstVo := reflect.ValueOf(dstPtr)
	if dstVo.Kind() != reflect.Ptr {
		return errors.New("rpclient: 'dstPtr' param value type must be a pointer")
	}

	// 检查指针是否为 nil
	if dstVo.IsNil() {
		return errors.New("rpclient: 'dstPtr' pointer cannot be nil")
	}

	// 判断来源和目的数据类型是否可转换
	// map <=> struct (互转)
	// slice <=> array (互转)
	// 其他类型必须严格匹配
	// 只在类型完全不兼容时才会报错
	srcType := reflect.TypeOf(r.Data)
	dstType := dstVo.Elem().Type()
	srcKind := srcType.Kind()
	dstKind := dstType.Kind()

	// 更完善的类型兼容性检查
	if !isTypeCompatible(srcKind, dstKind) {
		return fmt.Errorf("rpclient: %s cannot be converted to %s", srcKind, dstKind)
	}

	b, err := json.Marshal(r.Data)
	if err != nil {
		return fmt.Errorf("rpclient: failed to marshal data: %w", err)
	}

	if err = json.Unmarshal(b, dstPtr); err != nil {
		return fmt.Errorf("rpclient: failed to unmarshal data: %w", err)
	}

	return nil
}

// isTypeCompatible 检查源类型和目标类型是否兼容
func isTypeCompatible(srcKind, dstKind reflect.Kind) bool {
	// 相同类型总是兼容的
	if srcKind == dstKind {
		return true
	}

	switch srcKind {
	case reflect.Map, reflect.Struct:
		// map 和 struct 可以互转
		return dstKind == reflect.Map || dstKind == reflect.Struct
	case reflect.Slice, reflect.Array:
		// slice 和 array 可以互转
		return dstKind == reflect.Slice || dstKind == reflect.Array
	default:
		// 其他类型必须严格匹配
		return false
	}
}
