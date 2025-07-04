package rpclient

import (
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

type Args []*Payload

type RpcClient struct {
	*rpc.Client
	logger *slog.Logger
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
	}))
	conn, err := net.Dial(opt.Network, addr)
	if err != nil {
		logger.Error(err.Error())
		return &RpcClient{}, rrse.E(rrse.Op("dial"), err)
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
	c.logger.WithGroup("rpc").Info("call", "serviceMethod", serviceMethod, "args", *&args, "reply", reply, "error", err)
	return err
}

// Close closes both the network connection and the RPC client.
// It appends any errors encountered during the closing of the connection
// or the client to the Error field. If the connection or client is nil,
// it skips the closing operation for that component.
func (c *RpcClient) Close() error {
	if c.Client != nil {
		if err := c.Client.Close(); err != nil {
			err = rrse.E(rrse.Op("close"), err)
			c.logger.WithGroup("rpc").Error("close", "error", err)
			return err
		}
	}
	return nil
}
