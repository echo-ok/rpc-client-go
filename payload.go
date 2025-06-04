package client

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
	if len(body) == 0 {
		p.Body = nil
	} else if len(body) == 1 {
		p.Body = body[0]
	} else {
		p.Body = body
	}

	return p
}
