package rpclient

import (
	"reflect"
	"strings"
)

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
