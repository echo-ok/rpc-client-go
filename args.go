package rpclient

import (
	"reflect"
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

	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Slice, reflect.Map, reflect.Array:
		return val.Len() == 0
	case reflect.String:
		return val.Len() == 0 || val.String() == ""
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
