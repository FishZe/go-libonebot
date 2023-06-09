package go_libonebot

import (
	"github.com/FishZe/go-libonebot/protocol"
	"sync"
)

// ActionMux 请求处理接口
type ActionMux struct {
	actionHandleFunc sync.Map
}

// NewActionMux 获取一个ActionMux
func NewActionMux() (a *ActionMux) {
	a = new(ActionMux)
	return
}

// AddRequestInterface 添加一个请求处理接口
//
// action: 请求的action名称
//
// f: 请求处理接口 protocol.RequestInterface 的实现
func (a *ActionMux) AddRequestInterface(action string, f protocol.RequestInterface) {
	a.actionHandleFunc.Store(action, f)
}

// DeleteRequestInterface 删除一个请求处理接口
func (a *ActionMux) DeleteRequestInterface(action string) {
	a.actionHandleFunc.Delete(action)
}

// GetRequestInterface 获取一个请求处理接口
func (a *ActionMux) GetRequestInterface(action string) (f protocol.RequestInterface) {
	ff, ok := a.actionHandleFunc.Load(action)
	if !ok {
		return nil
	}
	return ff.(protocol.RequestInterface)
}
