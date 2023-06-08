package go_libonebot

import (
	"errors"
	"github.com/FishZe/go-libonebot/connect"
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
)

var (
	// ErrorConnectionUUIDNotFound 连接UUID未找到
	ErrorConnectionUUIDNotFound = errors.New("connection uuid not found")
)

// sendResponse 发送一个响应
func (b *Bot) sendResponse(e any, request *protocol.Request) error {
	// 检查Response是否合法
	// 要返回客户端要求的Echo
	err := protocol.ResponseCheck(request, e)
	if err != nil {
		// 不是合法的响应 返回一个20001 bad handler 错误
		util.Logger.Error("response check error: " + err.Error())
		_ = b.sendResponse(protocol.NewEmptyResponse(protocol.ResponseCodeBadHandler), request)
		return err
	}
	// 通过连接UUID发送
	if c, ok := b.connections.Load(request.ConnectionUUID); ok {
		if err = (*(c.(*connect.Connection))).ConnectSendResponse(e); err != nil {
			return err
		}
		return nil
	}
	// 未找到连接 返回错误
	return ErrorConnectionUUIDNotFound
}
