package client

import (
	"github.com/spf13/cast"
)

type Configuration map[string]any

func (c Configuration) GetString(key string) string {
	return cast.ToString(c[key])
}

type Store struct {
	ID            string        `json:"id"`            // 店铺 ID
	Name          string        `json:"name"`          // 店铺名称
	Env           string        `json:"env"`           // 运行环境
	Debug         bool          `json:"debug"`         // 是否开启调试模式
	Configuration Configuration `json:"configuration"` // 配置
}
