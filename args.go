package rpclient

import (
	"reflect"
	"strings"
)

// Args 查询参数
type Args []*Payload

func NewArgs() Args {
	return Args{}
}

// isEmpty 判断变量是否为空
// nil, "", []T{}, [...]T{}, map[T]V{} 均为空
func isEmpty(v any) bool {
	if v == nil {
		return true
	}

	vo := reflect.ValueOf(v)
	switch vo.Kind() {
	case reflect.Slice, reflect.Map, reflect.Array:
		return vo.Len() == 0
	case reflect.Ptr:
		val := vo.Elem()
		if val.Kind() == reflect.Array {
			return val.Len() == 0
		}
		return false
	case reflect.String:
		return vo.Len() == 0 || strings.TrimSpace(vo.String()) == ""
	default:
		return false
	}
}

// IsEmpty 是否为空
func (a Args) IsEmpty() bool {
	return len(a) == 0
}

// Add 添加查询
func (a Args) Add(payload *Payload) Args {
	for _, v := range a {
		// 同一个店铺，且参数一致的情况下忽略掉
		v1 := v.Body
		v2 := payload.Body
		if v.Store.ID == payload.Store.ID && ((isEmpty(v1) && isEmpty(v2)) || reflect.DeepEqual(v1, v2)) {
			return a
		}
	}
	return append(a, payload)
}

// Del 删除查询
func (a Args) Del(storeId string) Args {
	aa := Args{}
	for _, v := range a {
		if v.Store.ID == storeId {
			continue
		}
		aa = append(aa, v)
	}
	return aa
}

// SetBody 设置所有存储的查询参数
func (a Args) SetBody(body any) Args {
	aa := Args{}
	for _, v := range a {
		v.Body = body
		aa = append(aa, v)
	}
	return aa
}

// SetStoreBody 设置指定存储的查询参数
func (a Args) SetStoreBody(storeId string, body any) Args {
	aa := Args{}
	for _, v := range a {
		if v.Store.ID == storeId {
			v.Body = body
		}
		aa = append(aa, v)
	}
	return aa
}
