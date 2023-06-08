## 快速上手

### 创建配置

首先, 您需要在项目中创建一个配置, 包含您的实现的基础信息, 例如:
```go
import (
    onebot "go-libonebot"
)

//...

config := onebot.OneBotConfig{PlatForm: "QQ", Version: "1.0.0", Implementation: "MyQQClient"}
```

其中, `PlatForm`表示您的实现的平台, `Version`表示您的实现的版本, `Implementation`表示您的实现的名称.

每个 OneBot 实现应有一个格式为 [a-z][\-a-z0-9]*(\.[\-a-z0-9]+)* 的名字，点号分割不同部分可以允许实现采用版本号来区分扩展 API 版本，或者方便兼容实现声明自己的兼容性。

### 创建机器人实例

将上述的`config`和`userId`传入`NewOneBot`函数, 即可创建一个机器人实例:

```go
userId := "123456"
bot := onebot.NewOneBot(config, userId)
```

### 创建连接

连接分为`V12`和`V11`两种， 每种都支持`http`, `webhook`, `websocket`, `websocket_reverse`四种, 现阶段只支持了`V12`, `V11`正在积极适配中

每个连接可以选择一种连接方式, 需要多种连接的, 可以创建多种连接
```go
v12 "go-libonebot/connect/v12"

//...

conn, err := v12.NewOneBotV12Connect(v12.OneBotV12Config{
    /*
        // http连接
        ConnectType: v12.ConnectTypeHttp,
        HttpConfig: v12.OneBotV12HttpConfig{
            Host: "127.0.0.1",
            Port: 20003,
            AccessToken: "123456",
        },
    */

    /*
        // Websocket连接
        ConnectType: v12.ConnectTypeWebSocket,
        WebsocketConfig: v12.OneBotV12WebsocketConfig{
            Host: "127.0.0.1",
            Port: 20003,
            AccessToken: "123456",
        },
    */
    /*
        // Webhook连接
        ConnectType: v12.ConnectTypeWebhook,
		WebhookConfig: v12.OneBotV12WebHookConfig{
			Url:         "http://127.0.0.1:8080/",
			AccessToken: "123456",
			TimeOut:     5000,
			UserAgent: "go-libonebot",
		},
    */
    // Websocket Reverse 连接
    ConnectType: v12.ConnectTypeWebSocketReverse,
    WebsocketReverseConfig: v12.OneBotV12WebsocketReverseConfig{
        Url:               "ws://192.168.81.137:20001",
        ReconnectInterval: 5000,
        AccessToken: "123456",
    },
})
```

然后, 把连接加入bot实例中:
```go
err = b.AddConnection(conn)
```

`V12`协议支持多机器人连接, 即同一种连接方式实现多个机器人, GoLibOneBot也支持了特性, 你可以一个机器人加入多个连接, 一个连接加入多个机器人.

### 上报事件

```go
import "go-libonebot/protocol"

//...

bot.SendEvent(someEvent)
```

事件分为`Meta`, `message`, `notice`, `request`, 目前所有的`v12`协议的事件都已内置, 可以直接使用, 例如:
```go
// 创建新的私聊事件
evt := protocol.NewMessageEventPrivate()
// 加入消息段
evt.Message = append(e.Message, protocol.GetSegmentText("今天吃什么"))
evt.UserId = "123456"
err = bot.SendEvent(evt)
```

除了已经内置的`Event`, 你也可以自己实现一个事件, 格式为:
```go
struct {
    *protocol.Event
    // 你的字段
}
```
在发送时, 所有的导出字段都将会`json`编码后作为`data`字段发送.

### 处理请求
    
```go
bot.AddRequestInterface(string, protocol.RequestInterface)
```
GoLibOneBot内置了所有的`v12`协议的请求, 你可以直接使用.

其中, `HandleActionxxx`参数为一个函数, 函数传入一个`RequestXXX`, 返回一个`ResponseXXX`, 例如:

`ResponseXXX`由NewResponseXXX构造 (请勿自行直接创建结构体), 传入返回的retcode, 表示是否正常执行.


```go
bot.AddRequestInterface(protocol.HandleActionSendMessage(func(e *protocol.RequestSendMessage) *protocol.ResponseSendMessage {
    msg := protocol.NewResponseSendMessage(0)
    msg.MessageId = util.GetUUID()
    return msg
}))
```

除了内置的`Action`, 你也可以自行实现, 首先要实现一个`ResponseXXX`:

```go

struct ResponseXXX {
    *protocol.Response
    // 你的字段
}


然后, 实现一个`RequestInterface`:

```go

type RequestInterfaceXXX struct {
    *protocol.Request
    // 你的字段
}

// New 构造函数
func (*RequestInterfaceXXX)New() any {
    return &RequestInterfaceXXX{
        Request: &Request{
			Action: "your_action_name",
		},
    }
}

// Do 处理函数
func (*RequestInterfaceXXX)Do() any {
    // 处理逻辑
    return &ResponseXXX{}
}

```