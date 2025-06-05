RPC Client
===========

## Install
```shell
go get -u github.com/echo-ok/rpc-client-go
```

## Usage
```go
func main() {
    var payload *Payload
    var reply Reply
	payload := NewPayload(Store{
        ID:    "-1",
        Name:  "Temu SEMI Store",
        Env:   "prod",
        Debug: true,
        Configuration: Configuration{
            "region":             "US",
            "app_key":            "app key",
            "app_secret":         "app secret",
            "access_token":       "access token",
            "static_file_server": "static file server address",
        },
	})
	rpcClient, err = NewRpcClient(cfg.RpcAddress, &Option{
		Network: "tcp", 
		Codec:   jsonCodec,
	})
    err := rpcClient.Call("Temu.Semi.Order.Query", Args{
		payload.SetBody(map[string]any{
			"parentOrderSnList": []string{"PO-211-19255520399990061"}, 
			"regionId":          211,
		}),
	}, &reply)
}
```