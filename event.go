package go_libonebot

import (
	"github.com/FishZe/go-libonebot/connect"
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
)

// SendEvent 发送一个事件
//
// e: 事件 需要包含*Event字段且满足EventCheck
func (b *Bot) SendEvent(e any) error {
	defer func() {
		// recover错误
		if err := recover(); err != nil {
			util.Logger.Error("handle event error: " + err.(error).Error())
		}
	}()
	// 检查Event是否合法
	eType, id, err := protocol.EventCheck(b.info.self, e)
	if err != nil {
		util.Logger.Warning("EventCheck failed when send event: " + err.Error())
		return err
	}
	// 对每一个连接发送
	b.connections.Range(func(key, value any) bool {
		// 为了不影响其他连接的上报, 不处理错误
		if value != nil {
			err = (*(value.(*connect.Connection))).ConnectSendEvent(e, id, eType)
			if err != nil {
				util.Logger.Debug("send event error: " + err.Error())
			}
		}
		return true
	})
	return nil
}
