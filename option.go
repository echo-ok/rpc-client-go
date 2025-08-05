package rpclient

const (
	JsonCodec    = "json"
	GoridgeCodec = "goridge"
)

// Option NetWork Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4" (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and "unixpacket".
// Codec supported codecs are "goridge" and "json"
type Option struct {
	Network        string   `json:"network" yaml:"network" toml:"network"`                         // Current only support `tcp`
	Codec          string   `json:"codec" yaml:"codec" toml:"codec"`                               // Codes: json, goridge
	LogLevel       string   `json:"log_level" yaml:"log_level" toml:"log_level"`                   // Level: debug, info, warn, error
	SensitiveWords []string `json:"sensitive_words" yaml:"sensitive_words" toml:"sensitive_words"` // Sensitive words
}

var defaultOption = Option{
	Network:  "tcp",
	Codec:    JsonCodec,
	LogLevel: "debug",
	SensitiveWords: []string{
		"app_key",
		"key",
		"app_secret",
		"secret",
		"access_token",
		"token",
		"username",
		"name",
		"password",
		"pwd",
	},
}
