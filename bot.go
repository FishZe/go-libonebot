package go_libonebot

import (
	"github.com/FishZe/go-libonebot/connect"
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
	"sync"
	"time"
)

// Bot 一个oneBot实例
type Bot struct {
	// connections 连接列表
	connections sync.Map
	// info 基本信息
	info struct {
		// self 自身字段 self
		self protocol.Self
		// version 版本
		version string
		// impl 实现名称
		impl string
	}
	// requestChan 请求通道
	requestChan chan protocol.RawRequestType
	// handleMux 请求处理接口
	handleMux *ActionMux
	// heartBeatInterval 心跳间隔
	heartBeatInterval int
}

// NewOneBot 获取一个OneBot实例
//
// 需要传入config配置和userId
func NewOneBot(config protocol.OneBotConfig, userId string) (b *Bot) {
	b = new(Bot)
	b.info.version = config.Version
	b.info.impl = config.Implementation
	b.info.self = protocol.Self{
		PlatForm: config.PlatForm,
		UserId:   userId,
	}
	if config.HeartBeat.Enable {
		if config.HeartBeat.Interval <= 0 {
			b.heartBeatInterval = 5000
		} else {
			b.heartBeatInterval = config.HeartBeat.Interval
		}
	}
	b.requestChan = make(chan protocol.RawRequestType, 65535)
	b.startRequestChan()
	// 开启心跳
	go b.startHeartBeat()
	return
}

// AddConnection 添加一个连接
func (b *Bot) AddConnection(c connect.Connection) error {
	// 双向绑定 向连接添加机器人
	err := c.AddBot(b.info.impl, b.info.version, c.GetVersion(), b.info.self)
	if err != nil {
		return err
	}
	// 绑定请求通道
	err = c.AddBotRequestChan(b.info.self, b.requestChan)
	if err != nil {
		return err
	}
	// 加入所有连接
	b.connections.Store(c.GetUUID(), &c)
	return nil
}

// Handle 绑定处理器
func (b *Bot) Handle(mux *ActionMux) {
	b.handleMux = mux
}

// startHeartBeat 开启心跳
// @receiver b
func (b *Bot) startHeartBeat() {
	if b.heartBeatInterval <= 0 {
		return
	}
	for {
		e := protocol.NewMetaEventHeartbeat()
		e.Interval = b.heartBeatInterval
		if err := b.SendEvent(e); err != nil {
			util.Logger.Error("send heartbeat error: " + err.Error())
		}
		time.Sleep(time.Duration(b.heartBeatInterval) * time.Millisecond)
	}
}
