package go_libonebot

import (
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
	"github.com/jinzhu/copier"
)

// startRequestChan 开始请求通道
func (b *Bot) startRequestChan() {
	go func() {
		for {
			select {
			case request := <-b.requestChan:
				go func() {
					send := false
					defer func() {
						// recover 错误
						if err := recover(); err != nil {
							util.PrintGoroutineStack()
							util.Logger.Error("handle action error: " + err.(error).Error())
						}
						// 没发送, 重新发送
						if !send {
							res := protocol.NewEmptyResponse(protocol.ResponseCodeBadHandler)
							_ = b.sendResponse(res, request.Request)
						}
					}()
					var res any
					// Action为空, 返回错误: 无效的动作请求参数 参数缺失或参数类型错误
					if request.Request.Action == "" {
						util.Logger.Debug("request action is empty")
						r := protocol.NewEmptyResponse(protocol.ResponseCodeBadParam)
						r.Message = "request action is empty"
						res = r
					} else {
						x := b.handleMux.GetRequestInterface(request.Request.Action)
						if x == nil {
							// 找不到: 不支持的动作请求	OneBot 实现没有实现该动作
							util.Logger.Debug("request action not found: " + request.Request.Action)
							r := protocol.NewEmptyResponse(protocol.ResponseCodeUnsupportedAction)
							r.Message = "unsupported action: " + request.Request.Action
							res = r
						} else {
							util.Logger.Debug("request action found: " + request.Request.Action + " / " + request.Request.GetID() + "handling...")
							handleInterface := x.New()
							// 拷贝过来
							err := copier.Copy(handleInterface, x)
							if err != nil {
								res = protocol.NewEmptyResponse(protocol.ResponseCodeBadParam)
							} else {
								// 将RawRequestType转换为RequestInterface
								err = protocol.RequestAdapter(handleInterface, request)
								if err != nil {
									res = protocol.NewEmptyResponse(protocol.ResponseCodeBadParam)
								} else {
									res = handleInterface.(protocol.RequestInterface).Do()
									util.Logger.Debug("request action: " + request.Request.Action + " / " + request.Request.GetID() + " handled")
								}
							}
						}
					}
					util.Logger.Debug("request action: " + request.Request.Action)
					// 发送响应
					err := b.sendResponse(res, request.Request)
					if err != nil {
						util.Logger.Warning("send response failed: " + err.Error())
					}
					send = true
				}()
			}
		}
	}()
}
