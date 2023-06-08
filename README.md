# Go-LibOneBot

## 什么是OneBot?

来自[onebot.dev](https://onebot.dev/introduction.html)的定义:

> OneBot 是一个聊天机器人应用接口标准，旨在统一不同聊天平台上的机器人应用开发接口，使开发者只需编写一次业务逻辑代码即可应用到多种机器人平台。

OneBot协议是一个通用的机器人标准, 可以通过这套标准在机器人和各种框架间进行通信.

如[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)就是一个onebot实现, 可以通过这个实现在QQ上使用OneBot协议.

又如[nonebot](https://github.com/nonebot/nonebot2)就是一个onebot框架, 开发者可以通过这个框架编写你的机器人逻辑.

通过onebot协议, 可以使这两者连接起来, 实现在QQ上使用nonebot框架编写的机器人逻辑.

目前使用较多的OneBot协议有两个版本, 分别是OneBot11和OneBot12, 本项目以OneBot12为主要实现. OneBot11以兼容的方式实现 (正在适配...).

### OneBot协议内容?

事件: 机器人平台对框架上报的事件: 如收到消息, 消息撤回, 群成员增加或减少等

动作: 框架要求机器人平台执行的动作: 如发送消息, 撤回消息, 群组踢人等

动作请求: 即框架向机器人平台发送的动作请求

动作响应: 即机器人平台对框架发送的动作请求的响应 表示动作是否执行成功

消息段: 机器人平台对消息的格式化表示, 如纯文本, 图片, 视频, 音频, 文件, 链接, 表情等

连接: 框架和机器人平台之间的连接, 用于传输以上内容, OneBot协议包含了HTTP, Webhook, Websocket, 反向Websocket四种连接方式

### 什么是Libonebot?

来自[onebot.dev](https://onebot.dev/ecosystem.html)的定义:

> LibOneBot 指的是不同 OneBot 实现可以复用的部分，可以帮助 OneBot 实现者快速在新的聊天机器人平台实现 OneBot 标准

简单来说, LibOneBot已经帮你实现了基本的连接和消息处理, 你只需要实现你的业务逻辑, 就可以快速制作出一个OneBot实现,
做出一个像go-cqhttp这样的机器人客户端.

### 什么是Go-LibOneBot?

Go-LibOneBot 是一个go语言的LibOneBot, 可以通过这个库用go语言快速制作出一个OneBot实现.

内置了所有的OneBot12协议的事件和动作, 上手简单, 你只需要实现你的业务逻辑, 就可以快速制作出一个OneBot实现.

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
- [x] 自定义Json解析器和Logger

## 快速开始

```go
// 配置Bot
config := onebot.OneBotConfig{PlatForm: "QQ", Version: "1.0.0", Implementation: "go-cq"}
// 创建Bot
bot := onebot.NewOneBot(c, "123456")
// 创建v12连接
conn, _ := v12.NewOneBotV12Connect(v12.OneBotV12Config{
// 选择HTTP连接 
ConnectType: v12.ConnectTypeHttp,
HttpConfig: v12.OneBotV12HttpConfig{
// Host 地址
Host: "127.0.0.1",
// Port 端口
Port: 20003,
},

})
// 把这个连接添加到Bot中
_ = bot.AddConnection(conn)
// 添加一个事件处理器 : 发送消息请求
b.AddRequestInterface(protocol.HandleActionSendMessage(func(e *protocol.RequestSendMessage) *protocol.ResponseSendMessage {
log.Println("收到框架事件请求 发送消息: ", e.Message)
// 事件回复 状态码为0正常
msg := protocol.NewResponseSendMessage(0)
// 补充动作响应 填写MessageID
msg.MessageId = util.GetUUID()
return msg
}))
{
// 上报收到私聊消息事件
evt := protocol.NewMessageEventPrivate()
// 加入消息段
evt.Message = append(evt.Message, protocol.GetSegmentText("你好啊"))
// 设置用户id
evt.UserId = "123456"
_ = bot.SendEvent(evt)
}
```

具体使用方式请参考: [使用文档](doc/README.md)