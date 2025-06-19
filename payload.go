package rpclient

type Payload struct {
	Store Store `json:"store"`
	Body  any   `json:"body"`
}

func NewPayload(store Store) *Payload {
	return &Payload{
		Store: store,
	}
}

func (p *Payload) SetBody(body ...any) *Payload {
	switch len(body) {
	case 0:
		p.Body = nil
	case 1:
		p.Body = body[0]
	default:
		p.Body = body
	}

	return p
}
