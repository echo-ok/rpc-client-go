package rpclient

import (
	"log/slog"
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
		Network:  "tcp",
		Codec:    jsonCodec,
		LogLevel: slog.LevelDebug,
		Debug:    true,
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
			Items []any `json:"items"`
		}
		err = result.ConvertDataTo(&pager)
		assert.NoError(t, err)
		assert.Equal(t, pager.TotalCount, len(pager.Items))
	}
}
