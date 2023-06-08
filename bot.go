package go_libonebot

import (
	"github.com/FishZe/go-libonebot/connect"
	"github.com/FishZe/go-libonebot/protocol"
	"sync"
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
	// requestHandleFunc 请求处理函数
	requestHandleFunc sync.Map
}

// NewOneBot 获取一个OneBot实例
//
// 需要传入config配置和userId
func NewOneBot(config OneBotConfig, userId string) (b *Bot) {
	b = new(Bot)
	b.info.version = config.Version
	b.info.impl = config.Implementation
	b.info.self = protocol.Self{
		PlatForm: config.PlatForm,
		UserId:   userId,
	}
	b.requestChan = make(chan protocol.RawRequestType, 65535)
	b.startRequestChan()
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
