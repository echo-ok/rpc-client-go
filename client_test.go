package client

import (
	"github.com/goccy/go-json"
	"github.com/hiscaler/temu-go/entity"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	payload   Payload
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

func init() {
	var cfg *config
	b, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(b, &cfg); err != nil {
		panic(err)
	}

	payload = NewPayload("-1", "Temu SEMI Store", Configuration{
		"region":             cfg.Region,
		"app_key":            cfg.AppKey,
		"app_secret":         cfg.AppSecret,
		"access_token":       cfg.AccessToken,
		"static_file_server": cfg.StaticFileServer,
	})
	rpcClient, err = NewRpcClient(cfg.RpcAddress, &Option{
		Network: "tcp",
		Codec:   jsonCodec,
	})
	if err != nil {
		panic(err)
	}
}

func TestTemuSemiOrderCustomizationInformation(t *testing.T) {
	reply.Reset()
	err := rpcClient.Call("Temu.Semi.Order.CustomizationInformation", []Payload{
		payload.SetBody([]string{"211-12297657592950317"}),
	}, &reply)
	assert.NoError(t, err)
	for _, result := range reply.Results {
		if !result.Ok {
			continue
		}

		var infos []entity.SemiOrderCustomizationInformation
		err = result.ConvertDataTo(&infos)
		assert.NoError(t, err)
	}
}

func TestTemuSemiOrder(t *testing.T) {
	reply.Reset()
	err := rpcClient.Call("Temu.Semi.Order.Query", []Payload{
		payload.
			SetBody(map[string]any{
				"parentOrderSnList": []string{"PO-211-19255520399990061"},
				"regionId":          211,
			}),
	}, &reply)
	assert.NoError(t, err)
	for _, result := range reply.Results {
		if !result.Ok {
			continue
		}

		var value struct {
			Items []entity.ParentOrder `json:"items"`
		}
		err = result.ConvertDataTo(&value)
		assert.NoError(t, err)
	}
}
