package rpclient

import (
	"os"
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

var (
	payload   *Payload
	rpcClient *RpcClient
	reply     Reply
)

type config struct {
	RpcAddress       string `json:"rpc_address"`
	Region           string `json:"region"`
	AppKey           string `json:"app_key"`
	AppSecret        string `json:"app_secret"`
	AccessToken      string `json:"access_token"`
	StaticFileServer string `json:"static_file_server"`
}

func TestMain(m *testing.M) {
	var cfg *config
	b, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(b, &cfg); err != nil {
		panic(err)
	}

	payload = NewPayload(Store{
		ID:    "-1",
		Name:  "Temu SEMI Store",
		Env:   "prod",
		Debug: true,
		Configuration: Configuration{
			"region":             cfg.Region,
			"app_key":            cfg.AppKey,
			"app_secret":         cfg.AppSecret,
			"access_token":       cfg.AccessToken,
			"static_file_server": cfg.StaticFileServer,
		},
	})
	rpcClient, err = NewClient(cfg.RpcAddress, &Option{
		Network:        "tcp",
		Codec:          JsonCodec,
		LogLevel:       "debug",
		SensitiveWords: []string{"access_token", "app_key", "app_secret"},
	})
	if err != nil {
		panic(err)
	}
	defer func(rpcClient *RpcClient) {
		_ = rpcClient.Close()
	}(rpcClient)
	m.Run()
}

func TestTemuSemiOrderCustomizationInformation(t *testing.T) {
	err := rpcClient.Call("Temu.Semi.Order.CustomizationInformation", NewArgs().Add(payload.SetBody([]string{"211-12297657592950317"})), &reply)
	assert.NoError(t, err)
	assert.Equal(t, len(reply.Results), 1)
	for _, result := range reply.Results {
		if !result.Ok {
			continue
		}

		var infos []any
		err = result.ConvertDataTo(&infos)
		assert.NoError(t, err)
	}
}

func TestTemuSemiOrder(t *testing.T) {
	reply.Reset()
	err := rpcClient.Call("Temu.Semi.Order.Query", NewArgs().Add(payload.SetBody(map[string]any{
		"parentOrderSnList": []string{"PO-211-19255520399990061"},
		"regionId":          211,
	})), &reply)
	assert.NoError(t, err)
	assert.Equal(t, len(reply.Results), 1)
	for _, result := range reply.Results {
		if !result.Ok {
			continue
		}

		var pager struct {
			Pager
			Items []Result `json:"items"`
		}
		err = result.ConvertDataTo(&pager)
		assert.NoError(t, err)
		assert.Equal(t, pager.TotalCount, len(pager.Items))
	}
}

func TestTemuGoods(t *testing.T) {
	reply.Reset()
	err := rpcClient.Call("Temu.Goods.Detail", NewArgs().Add(payload.SetBody(11)), &reply)
	assert.NoError(t, err)
	assert.Equal(t, len(reply.Results), 1)
	for _, result := range reply.Results {
		if !result.Ok {
			continue
		}

		var data Result
		err = result.ConvertDataTo(&data)
		assert.NoError(t, err)
	}
}

// 半托订单发货信息
func TestTemuSemiOrderShippingInformation(t *testing.T) {
	reply.Reset()
	err := rpcClient.Call("Temu.Semi.Order.ShippingInformation", NewArgs().Add(payload.SetBody("PO-211-12969515438712454")), &reply)
	assert.NoError(t, err)
	assert.Equal(t, len(reply.Results), 1)
	for _, result := range reply.Results {
		if !result.Ok {
			continue
		}

		var data any
		err = result.ConvertDataTo(&data)
		assert.NoError(t, err)
	}
}

// 半托订单定制信息
func TestTemuSemiCustomizationInformation(t *testing.T) {
	reply.Reset()
	err := rpcClient.Call("Temu.Semi.Order.CustomizationInformation", NewArgs().Add(payload.SetBody([]string{
		"211-12969609810552454",
		"211-12969567867512454",
	})), &reply)
	assert.NoError(t, err)
	assert.Equal(t, len(reply.Results), 1)
	for _, result := range reply.Results {
		if !result.Ok {
			continue
		}

		var data any
		err = result.ConvertDataTo(&data)
		assert.NoError(t, err)
	}
}
