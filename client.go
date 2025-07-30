package rpclient

import (
	"fmt"
	"log/slog"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"

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
	debug  bool
	logger *slog.Logger
	dsn    string
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
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: opt.LogLevel,
	})).WithGroup("rpclient")
	dsn := fmt.Sprintf("%s://%s", opt.Network, addr)
	conn, err := net.Dial(opt.Network, addr)
	if err != nil {
		logger.Error("Dial", "dsn", dsn, "error", err)
		return nil, rrse.E(rrse.Op("dial"), err)
	}

	debug := opt.Debug
	if debug {
		logger.Info("Dial", "dsn", dsn, "error", nil)
	}
	var clientCodec rpc.ClientCodec
	if opt.Codec == goridgeCodec {
		clientCodec = goridgeRpc.NewClientCodec(conn)
	} else {
		clientCodec = jsonrpc.NewClientCodec(conn)
	}
	return &RpcClient{
		Client: rpc.NewClientWithCodec(clientCodec),
		logger: logger,
		debug:  debug,
		dsn:    dsn,
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
	loggerArgs := []any{"dsn", c.dsn, "serviceMethod", serviceMethod, "args", *&args, "reply", reply, "error", err}
	if err != nil {
		c.logger.Error("Call", loggerArgs...)
	} else if c.debug {
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
		c.logger.Error("Close", "dsn", c.dsn, "error", err)
		return err
	}
	if c.debug {
		c.logger.Info("Close", "dsn", c.dsn, "error", nil)
	}
	return nil
}
