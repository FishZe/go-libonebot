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
	conn, err := v12.NewOneBotV12Connect(v12.OneBotV12Config{
		// 连接类型
		// Http
		ConnectType: v12.ConnectTypeHttp,
		HttpConfig: v12.OneBotV12HttpConfig{
			Host:            "127.0.0.1",
			Port:            20003,
			EventEnable:     true,
			EventBufferSize: 500,
		},
		/*
			// Websocket
			ConnectType: v12.ConnectTypeWebSocket,
			WebsocketConfig: v12.OneBotV12WebsocketConfig{
				Host: "127.0.0.1",
				Port: 20003,
			},
		*/
		/*
			// Websocket Reverse
			ConnectType: v12.ConnectTypeWebSocketReverse,
			WebsocketReverseConfig: v12.OneBotV12WebsocketReverseConfig{
				Url:               "ws://192.168.81.137:20001",
				ReconnectInterval: 5000,
			},
		*/
		/*
			// Webhook
			ConnectType: v12.ConnectTypeWebhook,
			WebhookConfig: v12.OneBotV12WebHookConfig{
				Url:         "http://127.0.0.1:8080/",
				AccessToken: "123456",
				TimeOut:     5000,
				UserAgent:   "go-libonebot",
			},
		*/
	})
	if err != nil {
		panic(err)
	}
	// 把连接加入到bot
	err = bot.AddConnection(conn)
	if err != nil {
		panic(err)
	}
	conn2, err := v12.NewOneBotV12Connect(v12.OneBotV12Config{
		ConnectType: v12.ConnectTypeWebSocketReverse,
		WebsocketReverseConfig: v12.OneBotV12WebsocketReverseConfig{
			Url:               "ws://192.168.81.137:20001",
			ReconnectInterval: 5000,
		},
	})
	if err != nil {
		panic(err)
	}
	err = bot.AddConnection(conn2)
	if err != nil {
		panic(err)
	}
	conn3, err := v12.NewOneBotV12Connect(v12.OneBotV12Config{
		ConnectType: v12.ConnectTypeWebSocket,
		WebsocketConfig: v12.OneBotV12WebsocketConfig{
			Host: "127.0.0.1",
			Port: 20004,
		},
	})
	if err != nil {
		panic(err)
	}
	err = bot.AddConnection(conn3)
	if err != nil {
		panic(err)
	}
	// 创建一个事件处理器
	mux := onebot.NewActionMux()
	mux.AddRequestInterface(protocol.HandleActionSendMessage(func(e *protocol.RequestSendMessage) *protocol.ResponseSendMessage {
		// 处理发送消息动作
		log.Println("SendMessage: ", e.Message)
		util.Logger.Info("SendMessage: " + e.Message[0].Data["text"].(string))
		msg := protocol.NewResponseSendMessage(0)
		msg.MessageId = util.GetUUID()
		return msg
	}))

	/*
		// 自定义动作的实现
		mux.AddRequestInterface(&MyActionRequest{})
	*/

	// 把事件处理器加入到bot
	bot.Handle(mux)
	for {
		// 每20秒发送一个收到私聊信息的事件
		e := protocol.NewMessageEventPrivate()
		// 消息为 "今天吃什么"
		e.Message = append(e.Message, protocol.GetSegmentText("今天吃什么"))
		e.UserId = "1234567"
		err = bot.SendEvent(e)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(20 * time.Second)
	}
	/*
		// 自定义事件的实现
		type MyEvent struct {
			*protocol.Event
			// 自定义字段
			MyParam string `json:"my_param"`
		}
		bot.SendEvent(&MyEvent{
			Event: &protocol.Event{
				Type: "message",
				SubType: "my_message",
			},
			MyParam: "hello world!"
		})
	*/
}

// MyActionRequest
// 自定义一个动作
// 动作请求
type MyActionRequest struct {
	*protocol.Request
	// 自定义字段
	MyParam string `json:"my_param"`
}

// MyActionResponse
// 动作响应
type MyActionResponse struct {
	*protocol.Response
	// 自定义字段
	MyData string `json:"my_data"`
}

// New
// 动作请求构造函数 包含Action名称
func (*MyActionRequest) New() any {
	return &MyActionRequest{
		Request: &protocol.Request{
			Action: "my_action",
		},
	}
}

// Do
// 动作请求执行函数 返回动作响应
func (r *MyActionRequest) Do() any {
	return &MyActionResponse{
		Response: &protocol.Response{
			// RetCode 字段不可缺少, 表示是否成功执行
			Retcode: protocol.ResponseCodeOk,
		},
		MyData: "Hello World",
	}
}
