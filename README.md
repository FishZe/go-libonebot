# Go-LibOneBot

## 什么是OneBot?

来自[onebot.dev](https://onebot.dev/introduction.html)的定义:

> `OneBot` 是一个聊天机器人应用接口标准，旨在统一不同聊天平台上的机器人应用开发接口，使开发者只需编写一次业务逻辑代码即可应用到多种机器人平台。

`OneBot`协议是一个通用的机器人标准, 可以通过这套标准在机器人和各种框架间进行通信.

如[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)就是一个onebot实现, 可以通过这个实现在QQ上使用OneBot协议.

又如[nonebot](https://github.com/nonebot/nonebot2)就是一个onebot框架, 开发者可以通过这个框架编写你的机器人逻辑.

通过`onebot`协议, 可以使这两者连接起来, 例如实现在`QQ`上使用`nonebot`框架编写的机器人逻辑.

目前使用较多的`OneBot`协议有两个版本, 分别是`OneBot11`和`OneBot12`, 本项目以`OneBot12`为主要实现. OneBot11以兼容的方式实现 (正在适配...).

### OneBot协议内容?

事件: 机器人平台对框架上报的事件: 如收到消息, 消息撤回, 群成员增加或减少等

动作: 框架要求机器人平台执行的动作: 如发送消息, 撤回消息, 群组踢人等

动作请求: 即框架向机器人平台发送的动作请求

动作响应: 即机器人平台对框架发送的动作请求的响应 表示动作是否执行成功

消息段: 机器人平台对消息的格式化表示, 如纯文本, 图片, 视频, 音频, 文件, 链接, 表情等

连接: 框架和机器人平台之间的连接, 用于传输以上内容, `OneBot`协议包含了`HTTP`, `Webhook`, `Websocket`, `反向Websocket`四种连接方式, 在本项目中，你也可以自行定义`OneBot`连接方式

### 什么是Libonebot?

来自[onebot.dev](https://onebot.dev/ecosystem.html)的定义:

> LibOneBot 指的是不同 OneBot 实现可以复用的部分，可以帮助 OneBot 实现者快速在新的聊天机器人平台实现 OneBot 标准

简单来说, LibOneBot已经帮你实现了基本的连接和消息处理, 你只需要实现你的业务逻辑, 就可以快速制作出一个OneBot实现,
做出一个像go-cqhttp这样的机器人客户端.

### 什么是Go-LibOneBot?

Go-LibOneBot 是一个`go`语言的`LibOneBot`, 可以通过这个库用`go`语言快速制作出一个`OneBot`实现.

内置了所有的`OneBot 12`协议的事件和动作, 上手简单, 你只需要实现你的业务逻辑, 就可以快速制作出一个`OneBot`实现.

## 支持 & TODO:

OneBot12 协议:

- [x] OneBot12 所有事件
- [x] OneBot12 所有动作 (请求和响应)
- [x] OneBot12 消息段
- [x] HTTP 连接
- [x] Webhook 连接
- [x] Websocket 连接
- [x] 反向Websocket 连接
- [x] 扩展Onebot12 协议: 事件 / 动作 / 消息段

OneBot11 协议:

- [ ] 通用转换器
- [ ] OneBot11 所有事件
- [ ] OneBot11 所有API
- [ ] OneBot11 CQ码
- [ ] HTTP 连接
- [ ] Webhook 连接
- [ ] Websocket 连接
- [ ] 反向Websocket 连接

开发者自行扩展:

- [x] 自定义事件
- [x] 自定义动作接口
- [x] 自定义协议和连接方式
- [x] 自定义`Json`解析器和`Logger`

## 示例代码

```go
package main

import (
	onebot "github.com/FishZe/go-libonebot"
	v12 "github.com/FishZe/go-libonebot/connect/v12"
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
	"log"
	"time"
)

func main() {
	// 创建一份配置
	onebotConfig := onebot.OneBotConfig{PlatForm: "QQ", Version: "1.0.0", Implementation: "MyQQImpl"}
	// 创建一个bot
	bot := onebot.NewOneBot(onebotConfig, "123456")
	// 创建一个连接
	conn, _ := v12.NewOneBotV12Connect(v12.OneBotV12Config{
		// 连接类型
		// Http
		ConnectType: v12.ConnectTypeHttp,
		HttpConfig: v12.OneBotV12HttpConfig{
			Host:            "127.0.0.1",
			Port:            20003,
			EventEnable:     true,
			EventBufferSize: 500,
		},
	})
	// 把连接加入到bot
	_ = bot.AddConnection(conn)
	// 创建一个事件处理器
	mux := onebot.NewActionMux()
	mux.AddRequestInterface(protocol.HandleActionSendMessage(func(e *protocol.RequestSendMessage) *protocol.ResponseSendMessage {
		// 处理发送消息动作
		log.Println("SendMessage: ", e.Message)
		msg := protocol.NewResponseSendMessage(0)
		msg.MessageId = util.GetUUID()
		return msg
	}))
	// 把事件处理器加入到bot
	bot.Handle(mux)
	for {
		// 每20秒发送一个收到私聊信息的事件
		e := protocol.NewMessageEventPrivate()
		// 消息为 "今天吃什么"
		e.Message = append(e.Message, protocol.GetSegmentText("今天吃什么"))
		e.UserId = "123456"
		err = bot.SendEvent(e)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(20 * time.Second)
	}
}
```

## 使用文档 

[使用文档](doc/README.md)
