package rpclient

import (
	"fmt"
	"log/slog"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"slices"
	"strings"

	rrse "github.com/roadrunner-server/errors"
	goridgeRpc "github.com/roadrunner-server/goridge/v3/pkg/rpc"
)

const (
	Dev  = "dev"  // 开发环境
	Test = "test" // 测试环境
	Prod = "prod" // 生产环境
)

type RpcClient struct {
	*rpc.Client
	logger *slog.Logger
	option *Option
}

func maskString(s string) string {
	n := len(s)
	switch n {
	case 0:
		return s
	case 1, 2:
		// 长度小于等于 2，直接返回相同长度的 *
		return strings.Repeat("*", n)
	case 3, 4, 5, 6:
		// 保留首尾字符，中间用 * 填充
		return s[:1] + strings.Repeat("*", n-2) + s[n-1:]
	default:
		// 长度大于 6 时，保留前 3 位和后 3 位，中间替换为 ******
		return s[:3] + "******" + s[n-3:]
	}
}

// NewClient creates a new RPC client to the given address.
//
// The address should be given in the format "host:port".
//
// The opt parameter is an optional Option pointer that can be used to specify
// the network type and codec to be used. If nil, the defaultOption is used.
//
// The method returns a pointer to a new RpcClient and an error. The RpcClient
// is a simple wrapper around the rpc.Client and net.Conn. The Error field of
// the RpcClient is set to the error returned by the underlying Close methods.
//
// The supported codecs are "goridge" and "json". The default codec is "json".
func NewClient(addr string, opt *Option) (*RpcClient, error) {
	if opt == nil {
		opt = &defaultOption
	}
	logLevel := slog.LevelDebug
	switch opt.LogLevel {
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	}
	logger := slog.
		New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: false,
			Level:     logLevel,
		})).
		WithGroup("rpclient").
		With("dsn", fmt.Sprintf("%s://%s", opt.Network, addr))
	conn, err := net.Dial(opt.Network, addr)
	if err != nil {
		logger.Error("Dial", "error", err)
		return nil, rrse.E(rrse.Op("dial"), err)
	}

	logger.Debug("Dial", "error", nil)
	var clientCodec rpc.ClientCodec
	if opt.Codec == GoridgeCodec {
		clientCodec = goridgeRpc.NewClientCodec(conn)
	} else {
		clientCodec = jsonrpc.NewClientCodec(conn)
	}
	logger.Debug("NewClientCodec", "codec", opt.Codec, "error", nil)
	return &RpcClient{
		Client: rpc.NewClientWithCodec(clientCodec),
		logger: logger,
		option: opt,
	}, nil
}

// Call calls the RPC server with the given service method and arguments.
// It returns an error if the call fails.
func (c *RpcClient) Call(serviceMethod string, args Args, reply *Reply) error {
	reply.Reset()
	err := c.Client.Call(serviceMethod, args, reply)
	if err != nil {
		err = rrse.E(rrse.Op("call"), err)
	}

	sanitizedArgs := make([]Payload, len(args))
	for i, arg := range args {
		cfg := make(Configuration, len(arg.Store.Configuration))
		for key, value := range arg.Store.Configuration {
			if slices.Index(c.option.SensitiveWords, key) != -1 {
				switch value.(type) {
				case string:
					str, _ := value.(string)
					value = maskString(str)
				case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
					value = maskString(fmt.Sprintf("%d", value))
				}
			}
			cfg[key] = value
		}
		sanitizedArg := Payload{
			Store: arg.Store,
			Body:  arg.Body,
		}
		sanitizedArg.Store.Configuration = cfg
		sanitizedArgs[i] = sanitizedArg
	}
	loggerArgs := []any{"serviceMethod", serviceMethod, "args", sanitizedArgs, "reply", reply, "error", err}
	if err != nil {
		c.logger.Error("Call", loggerArgs...)
	} else {
		c.logger.Info("Call", loggerArgs...)
	}
	return err
}

// Close closes both the network connection and the RPC client.
// It appends any errors encountered during the closing of the connection
// or the client to the Error field. If the connection or client is nil,
// it skips the closing operation for that component.
func (c *RpcClient) Close() error {
	if c.Client == nil {
		return nil
	}

	if err := c.Client.Close(); err != nil {
		err = rrse.E(rrse.Op("close"), err)
		c.logger.Error("Close", "error", err)
		return err
	}
	c.logger.Debug("Close", "error", nil)
	return nil
}
