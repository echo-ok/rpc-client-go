package client

type Payload struct {
	Store Store `json:"store"`
	Body  any   `json:"body"`
}

func NewPayload(shopId, shopName string, cfg Configuration) Payload {
	return Payload{
		Store: Store{
			ID:    shopId,
			Name:  shopName,
			Env:   "prod",
			Debug: true,
			Configuration: map[string]any{
				"region":             cfg.GetString("region"),
				"app_key":            cfg.GetString("app_key"),
				"app_secret":         cfg.GetString("app_secret"),
				"access_token":       cfg.GetString("access_token"),
				"static_file_server": cfg.GetString("static_file_server"),
			},
		},
	}
}

func (p Payload) SetBody(body ...any) Payload {
	if len(body) == 0 {
		p.Body = nil
	} else if len(body) == 1 {
		p.Body = body[0]
	} else {
		p.Body = body
	}

	return p
}
