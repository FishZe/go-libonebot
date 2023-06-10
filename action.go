package go_libonebot

import (
	"errors"
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
	"reflect"
	"sync"
)

var (
	// ErrorActionExist action已存在
	ErrorActionExist = errors.New("action already exist")
	// ErrorActionEmpty action为空
	ErrorActionEmpty = errors.New("action is empty")
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
func (a *ActionMux) AddRequestInterface(f protocol.RequestInterface) error {
	// 检查是否是指针 且 指向的是一个结构体
	if f != nil && reflect.TypeOf(f).Kind() != reflect.Ptr || reflect.TypeOf(f).Elem().Kind() != reflect.Struct {
		util.Logger.Warning("request mux: arg not a onebot request struct")
		return protocol.ErrorInvalidRequest
	}
	// 查找*request头
	requestID := -1
	t := reflect.TypeOf(f).Elem()
	v := reflect.ValueOf(f).Elem()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == protocol.RequestType {
			requestID = i
			break
		}
	}
	// *request不存在
	if requestID == -1 {
		util.Logger.Warning("request mux: arg not a onebot request struct")
		return protocol.ErrorInvalidRequest
	}
	// 获取*request
	request := v.Field(requestID).Interface().(*protocol.Request)
	if request == nil {
		util.Logger.Warning("request mux: arg not a onebot request struct")
		return protocol.ErrorInvalidRequest
	}
	// 没有action
	if request.Action == "" {
		util.Logger.Warning("request mux: request action is empty")
		return ErrorActionEmpty
	}
	// action已存在
	if _, ok := a.actionHandleFunc.Load(request.Action); ok {
		util.Logger.Warning("request mux: request action already exist")
		return ErrorActionExist
	}
	a.actionHandleFunc.Store(request.Action, f)
	return nil
}

// DeleteRequestInterface 删除一个请求处理接口
func (a *ActionMux) DeleteRequestInterface(action string) {
	a.actionHandleFunc.Delete(action)
}

// GetRequestInterface 获取一个请求处理接口
func (a *ActionMux) GetRequestInterface(action string) (f protocol.RequestInterface) {
	util.Logger.Debug("mux action request: " + action)
	ff, ok := a.actionHandleFunc.Load(action)
	if !ok {
		util.Logger.Warning("mux action request: " + action + " not found")
		return nil
	}
	return ff.(protocol.RequestInterface)
}
