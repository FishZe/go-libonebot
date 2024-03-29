package connect

import "github.com/FishZe/go-libonebot/protocol"

// Connection connection接口 可自定义
type Connection interface {
	// SendEvent 发送事件
	SendEvent(e any, eId string, eType string) error
	// SendResponse 发送响应
	SendResponse(e any) error
	// AddBot 添加了一个Bot
	AddBot(impl string, version string, oneBotVersion string, self protocol.Self) error
	// AddBotRequestChan 添加了一个Bot的请求通道
	AddBotRequestChan(self protocol.Self, botRequestChan chan protocol.RawRequestType) error
	// GetVersion 返回连接的版本
	GetVersion() string
	// GetUUID 返回连接的UUID
	GetUUID() string
}
