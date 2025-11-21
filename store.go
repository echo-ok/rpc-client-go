package rpclient

import (
	"github.com/spf13/cast"
)

// Configuration 配置信息
type Configuration map[string]any

func (c Configuration) read(key string) any {
	return c[key]
}

func (c Configuration) GetString(key string) string {
	return cast.ToString(c.read(key))
}

// Set 设置配置项，如果已存在则覆盖
func (c Configuration) Set(key string, value any) Configuration {
	c[key] = value
	return c
}

type Store struct {
	ID            string        `json:"id"`            // 店铺 ID
	Name          string        `json:"name"`          // 店铺名称
	Env           string        `json:"env"`           // 运行环境（dev, test, prod）
	Debug         bool          `json:"debug"`         // 是否开启调试模式
	Timeout       int           `json:"timeout"`       // 请求超时时间
	Configuration Configuration `json:"configuration"` // 配置信息
}
