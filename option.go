package rpclient

const (
	jsonCodec    = "json"
	goridgeCodec = "goridge"
)

// Option NetWork Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4" (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and "unixpacket".
// Codec supported codecs are "goridge" and "json"
type Option struct {
	Network  string `json:"network" yaml:"network" toml:"network"`
	Codec    string `json:"codec" yaml:"codec" toml:"codec"`
	LogLevel string `json:"log_level" yaml:"log_level" toml:"log_level"` // deubg, info, warn, error
}

var defaultOption = Option{
	Network:  "tcp",
	Codec:    jsonCodec,
	LogLevel: "debug",
}
