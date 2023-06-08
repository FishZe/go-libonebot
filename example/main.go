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
	c := onebot.OneBotConfig{PlatForm: "QQ", Version: "1.0.0", Implementation: "go-cq"}
	b := onebot.NewOneBot(c, "123456")
	conn, err := v12.NewOneBotV12Connect(v12.OneBotV12Config{
		/*
			ConnectType: v12.ConnectTypeHttp,
			HttpConfig: v12.OneBotV12HttpConfig{
				Host: "127.0.0.1",
				Port: 20003,
			},
		*/

		/*
				ConnectType: v12.ConnectTypeWebSocket,

			WebsocketConfig: v12.OneBotV12WebsocketConfig{
				Host: "127.0.0.1",
				Port: 20003,
			},
		*/
		/*
			ConnectType: v12.ConnectTypeWebSocketReverse,
			WebsocketReverseConfig: v12.OneBotV12WebsocketReverseConfig{
				Url:               "ws://192.168.81.137:20001",
				ReconnectInterval: 5000,
			},*/
		ConnectType: v12.ConnectTypeWebhook,
		WebhookConfig: v12.OneBotV12WebHookConfig{
			Url:         "http://127.0.0.1:8080/",
			AccessToken: "123456",
			TimeOut:     5000,
			UserAgent:   "go-libonebot",
		},
	})
	if err != nil {
		panic(err)
	}
	err = b.AddConnection(conn)
	if err != nil {
		panic(err)
	}

	b.AddRequestInterface(protocol.HandleActionGetLatestEvents(func(e *protocol.RequestGetLatestEvents) *protocol.ResponseGetLatestEvents {
		res := protocol.NewResponseGetLatestEvents(0)
		h := protocol.NewMetaEventHeartbeat()
		res.Events = append(res.Events, h)
		log.Println("GetLatestEvents")
		return res
	}))
	b.AddRequestInterface(protocol.HandleActionSendMessage(func(e *protocol.RequestSendMessage) *protocol.ResponseSendMessage {
		log.Println("SendMessage")
		log.Println(e.Message)
		msg := protocol.NewResponseSendMessage(0)
		msg.MessageId = util.GetUUID()
		return msg
	}))
	for {
		{
			e := protocol.NewMessageEventPrivate()
			e.Message = append(e.Message, protocol.GetSegmentText("今天吃什么"))
			e.UserId = "123456"
			log.Println(e)
			err = b.SendEvent(e)
			if err != nil {
				log.Println(err)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
