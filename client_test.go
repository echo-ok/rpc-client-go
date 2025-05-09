package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRpcClient(t *testing.T) {
	rpcClient, err := NewRpcClient("127.0.0.1:6001", &Option{
		Network: "tcp",
		Codec:   "json",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer rpcClient.Close()
	var reply Reply
	err = rpcClient.Call("Temu.Semi.Order.CustomizationInformation", []Payload{
		NewPayload("-1", "Temu SEMI Store", Configuration{
			"region":             "US",
			"app_key":            "",
			"app_secret":         "",
			"access_token":       "",
			"static_file_server": "http://127.0.0.1:6002",
		}).
			SetBody([]string{"PO-211-05235016115888888"}),
	}, &reply)
	assert.NoError(t, err)
}
