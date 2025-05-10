package client

import (
	"errors"
	rrse "github.com/roadrunner-server/errors"
	goridgeRpc "github.com/roadrunner-server/goridge/v3/pkg/rpc"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type RpcClient struct {
	net.Conn
	*rpc.Client
	Error error
}

// NewRpcClient creates a new RPC client to the given address.
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
func NewRpcClient(addr string, opt *Option) (*RpcClient, error) {
	if opt == nil {
		opt = &defaultOption
	}
	conn, err := net.Dial(opt.Network, addr)
	if err != nil {
		return &RpcClient{}, rrse.E(rrse.Op("dial"), err)
	}

	var clientCodec rpc.ClientCodec
	if opt.Codec == "goridge" {
		clientCodec = goridgeRpc.NewClientCodec(conn)
	} else {
		clientCodec = jsonrpc.NewClientCodec(conn)
	}

	return &RpcClient{
		Conn:   conn,
		Client: rpc.NewClientWithCodec(clientCodec),
	}, nil
}

// Close closes both the network connection and the RPC client.
// It appends any errors encountered during the closing of the connection
// or the client to the Error field. If the connection or client is nil,
// it skips the closing operation for that component.
func (c *RpcClient) Close() {
	if c.Conn != nil {
		if err := c.Conn.Close(); err != nil {
			c.Error = errors.Join(c.Error, rrse.E(rrse.Op("close"), err))
		}
	}
	if c.Client != nil {
		if err := c.Client.Close(); err != nil {
			c.Error = errors.Join(c.Error, rrse.E(rrse.Op("close"), err))
		}
	}
}
