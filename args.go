package rpclient

import "reflect"

type Args []*Payload

func NewArgs() Args {
	return Args{}
}

func (a Args) Add(payload *Payload) Args {
	for _, v := range a {
		// 同一个店铺，且参数一致的情况下忽略掉
		if v.Store.ID == payload.Store.ID && (v.Body == nil && payload.Body == nil || reflect.DeepEqual(v.Body, payload.Body)) {
			return a
		}
	}
	return append(a, payload)
}
