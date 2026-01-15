# RPC Client

Go 语言实现的 RPC 客户端库，为 RPC 服务提供简化的方法调用支持。

## 功能特性

- **多编解码器支持**：支持 JSON 和 Goridge 两种编解码器
- **环境隔离**：支持 dev/test/prod 三种环境，自动切换配置
- **敏感数据保护**：自动在日志中脱敏敏感字段
- **错误聚合**：支持从多个结果中收集并汇总错误
- **类型安全转换**：提供安全的数据类型转换方法
- **分页支持**：内置分页数据结构
- **请求去重**：自动去除重复请求

## 安装

```bash
go get -u github.com/echo-ok/rpc-client-go
```

## 快速开始

```go
package main

import (
	"fmt"
	"log"

	"github.com/echo-ok/rpc-client-go"
)

func main() {
	payload := rpclient.NewPayload(rpclient.Store{
		ID:    "-1",
		Name:  "Temu SEMI Store",
		Env:   "prod",
		Debug: true,
		Configuration: rpclient.Configuration{
			"region":             "US",
			"app_key":            "your_app_key",
			"app_secret":         "your_app_secret",
			"access_token":       "your_access_token",
			"static_file_server": "http://your-server:6002",
		},
	})

	rpcClient, err := rpclient.NewClient("127.0.0.1:6001", &rpclient.Option{
		Network:  "tcp",
		Codec:    rpclient.JsonCodec,
		LogLevel: "debug",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer rpcClient.Close()

	var reply rpclient.Reply
	err = rpcClient.Call("Temu.Semi.Order.Query", rpclient.NewArgs().Add(
		payload.SetBody(map[string]any{
			"parentOrderSnList": []string{"PO-211-19255520399990061"},
			"regionId":          211,
		}),
	), &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
```

## 核心概念

### Payload

请求负载，包含店铺配置和请求体数据。

```go
payload := rpclient.NewPayload(store)
payload.SetBody(requestData)
```

### Args

请求参数集合，支持自动去重。

```go
args := rpclient.NewArgs()
args = args.Add(payload1)
args = args.Add(payload2)
```

### Reply

RPC 响应，包含多个店铺的执行结果。

```go
reply.HasError()     // 检查是否有错误
reply.Errors()       // 获取错误列表
reply.ErrorSummary() // 获取错误摘要
```

### Result

单个店铺的执行结果，支持类型安全的数据转换。

```go
for _, result := range reply.Results {
	var data MyStruct
	err := result.ConvertDataTo(&data)
}
```

## 配置选项

### Option

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| Network | string | 否 | 网络类型，默认 `tcp` |
| Codec | string | 否 | 编解码器，支持 `json`、`goridge`，默认 `json` |
| LogLevel | string | 否 | 日志级别，支持 `debug`、`info`、`warn`、`error`，默认 `debug` |
| SensitiveWords | []string | 否 | 敏感词列表，日志中会自动脱敏 |

### Store

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| ID | string | 是 | 店铺 ID |
| Name | string | 是 | 店铺名称 |
| Env | string | 否 | 运行环境，支持 `dev`、`test`、`prod`，默认 `dev` |
| Debug | bool | 否 | 是否开启调试模式 |
| Timeout | int | 否 | 请求超时时间 |
| Configuration | Configuration | 是 | 店铺配置信息 |

### Configuration 常用配置项

| 配置项 | 说明 |
|--------|------|
| region | 区域代码，如 `US` |
| app_key | 应用密钥 |
| app_secret | 应用密钥 |
| access_token | 访问令牌 |
| static_file_server | 静态文件服务器地址 |

## 高级用法

### 批量请求多个店铺

```go
args := rpclient.NewArgs()
args.Add(rpclient.NewPayload(store1).SetBody(params1))
args.Add(rpclient.NewPayload(store2).SetBody(params2))

var reply rpclient.Reply
rpcClient.Call("Some.Service.Method", args, &reply)

if reply.HasError() {
	for _, err := range reply.Errors() {
		fmt.Println(err)
	}
}
```

### 使用 Goridge 编解码器

```go
rpcClient, err := rpclient.NewClient(addr, &rpclient.Option{
	Network: "tcp",
	Codec:   rpclient.GoridgeCodec,
})
```

### 敏感数据脱敏

默认会对以下字段进行脱敏处理：
- key, app_key, china_app_key
- secret, app_secret, china_app_secret
- token, access_token, china_access_token
- name, username, pwd, password

自定义敏感词：

```go
rpcClient, err := rpclient.NewClient(addr, &rpclient.Option{
	SensitiveWords: []string{"custom_field"},
})
```

### 分页数据处理

```go
type OrderList struct {
	 rpclient.Pager
	 Orders []Order `json:"items"`
}

var reply rpclient.Reply
rpcClient.Call("Order.List", args, &reply)

for _, result := range reply.Results {
	var orderList OrderList
	result.ConvertDataTo(&orderList)
	fmt.Printf("Page %d/%d, Total: %d\n",
		orderList.Page, orderList.PageCount, orderList.TotalCount)
}
```

## 错误处理

```go
var reply rpclient.Reply
err := rpcClient.Call("Service.Method", args, &reply)
if err != nil {
	// RPC 调用级别错误
	log.Fatal(err)
}

if reply.HasError() {
	// 业务级别错误
	for _, e := range reply.Errors() {
		log.Printf("Error: %v", e)
	}
}
```

## 测试

```bash
# 运行所有测试
go test -v ./...

# 运行指定测试文件
go test -v client_test.go

# 运行并查看覆盖率
go test -cover ./...
```

## 项目结构

```
├── client.go      # RPC 客户端核心实现
├── payload.go     # 请求负载结构
├── args.go        # 请求参数集合
├── reply.go       # 响应结构
├── result.go      # 结果结构
├── store.go       # 店铺配置
├── option.go      # 客户端配置
├── pager.go       # 分页数据结构
└── *_test.go      # 测试文件
```

## 许可证

MIT
