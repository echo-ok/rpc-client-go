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

// isEmptyValue 判断变量是否为空
// nil, "", []T{}, [...]T{}, map[T]V{} 均为空
func isEmptyValue(v any) bool {
	if v == nil {
		return true
	}

	vo := reflect.ValueOf(v)
	switch vo.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Slice, reflect.Map, reflect.Array:
		return vo.Len() == 0
	case reflect.Ptr:
		return isEmptyValue(vo.Elem().Interface())
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
		v1 := v.Body
		v2 := payload.Body
		if v.Store.ID == payload.Store.ID && ((isEmptyValue(v1) && isEmptyValue(v2)) || reflect.DeepEqual(v1, v2)) {
			// 同一个店铺，且参数一致的情况下忽略掉
			return append(Args{}, a...)
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
	aa := make(Args, len(a))
	for k, v := range a {
		p := &Payload{Store: v.Store}
		p.SetBody(body)
		aa[k] = p
	}
	return aa
}

// SetStoreBody 设置指定存储的查询参数
func (a Args) SetStoreBody(storeId string, body any) Args {
	if storeId == "" {
		return a
	}
	for k, v := range a {
		if v.Store.ID == storeId {
			a[k].SetBody(body)
		}
	}
	return a
}
