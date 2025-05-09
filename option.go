package client

type Option struct {
	Network string
	Codec   string
}

var defaultOption = Option{Network: "tcp", Codec: "json"}
