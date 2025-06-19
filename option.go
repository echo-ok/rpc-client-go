package rpclient

import "log/slog"

const (
	jsonCodec    = "json"
	goridgeCodec = "goridge"
)

// Option NetWork Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4" (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and "unixpacket".
// Codec supported codecs are "goridge" and "json"
type Option struct {
	Network  string
	Codec    string
	logLevel slog.Leveler
}

var defaultOption = Option{
	Network:  "tcp",
	Codec:    jsonCodec,
	logLevel: slog.LevelDebug,
}
