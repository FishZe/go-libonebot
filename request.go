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
					var res any
					// Action为空, 返回错误: 无效的动作请求参数 参数缺失或参数类型错误
					if request.Request.Action == "" {
						util.Logger.Debug("request action is empty")
						r := protocol.NewEmptyResponse(protocol.ResponseCodeBadParam)
						r.Message = "request action is empty"
						res = r
					} else {
						x, ok := b.requestHandleFunc.Load(request.Request.Action)
						if !ok {
							// 找不到: 不支持的动作请求	OneBot 实现没有实现该动作
							util.Logger.Debug("request action not found: " + request.Request.Action)
							r := protocol.NewEmptyResponse(protocol.ResponseCodeUnsupportedAction)
							r.Message = "unsupported action: " + request.Request.Action
							res = r
						} else {
							handleInterface := x.(protocol.RequestInterface).New()
							// 拷贝过来
							err := copier.Copy(handleInterface, x)
							if err != nil {
								res = protocol.NewEmptyResponse(protocol.ResponseCodeBadParam)
							} else {
								// 将RawRequestType转换为RequestInterface
								err = protocol.RawRequestTypeToRequest(handleInterface, request)
								if err != nil {
									res = protocol.NewEmptyResponse(protocol.ResponseCodeBadParam)
								} else {
									res = handleInterface.(protocol.RequestInterface).Do()
								}
							}
						}
					}
					util.Logger.Debug("request action: " + request.Request.Action)
					// 发送响应
					err := b.sendResponse(res, request.Request)
					if err != nil {
						util.Logger.Error("send response failed: " + err.Error())
					}
				}()
			}
		}
	}()
}

// AddRequestInterface 添加一个请求处理接口
//
// action: 请求的action名称
//
// f: 请求处理接口 protocol.RequestInterface 的实现
func (b *Bot) AddRequestInterface(action string, f protocol.RequestInterface) {
	b.requestHandleFunc.Store(action, f)
}
