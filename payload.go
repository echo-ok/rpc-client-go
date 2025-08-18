package rpclient

import (
	"slices"
	"strings"
)

type Payload struct {
	Store Store `json:"store"`
	Body  any   `json:"body"`
}

func NewPayload(store Store, body ...any) *Payload {
	store.Env = strings.ToLower(store.Env)
	if store.Env == "" || slices.Index([]string{Dev, Test, Prod}, store.Env) == -1 {
		store.Env = Dev
	}
	p := &Payload{
		Store: store,
	}
	p.SetBody(body...)
	return p
}

func (p *Payload) SetBody(body ...any) *Payload {
	if body == nil {
		p.Body = nil
		return p
	}

	switch len(body) {
	case 0:
		return p
	case 1:
		p.Body = body[0]
	default:
		p.Body = body
	}
	return p
}
