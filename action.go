package go_libonebot

import (
	"errors"
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
	"reflect"
	"regexp"
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
	regexHandleFunc  sync.Map
}

// regexActionMux 正则请求处理接口
type regexActionMux struct {
	Re *regexp.Regexp
	F  protocol.RequestInterface
}

// NewActionMux 获取一个ActionMux
func NewActionMux() (a *ActionMux) {
	a = new(ActionMux)
	return
}

func (a *ActionMux) checkRequestInterface(f protocol.RequestInterface) (*protocol.Request, error) {
	// 检查是否是指针 且 指向的是一个结构体
	if f != nil && reflect.TypeOf(f).Kind() != reflect.Ptr || reflect.TypeOf(f).Elem().Kind() != reflect.Struct {
		util.Logger.Warning("request mux: arg not a onebot request struct")
		return nil, protocol.ErrorInvalidRequest
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
		return nil, protocol.ErrorInvalidRequest
	}
	// 获取*request
	request := v.Field(requestID).Interface().(*protocol.Request)
	if request == nil {
		util.Logger.Warning("request mux: arg not a onebot request struct")
		return nil, protocol.ErrorInvalidRequest
	}
	// 没有action
	if request.Action == "" {
		util.Logger.Warning("request mux: request action is empty")
		return nil, ErrorActionEmpty
	}
	return request, nil
}

// AddRequestInterface 添加一个请求处理接口
//
// action: 请求的action名称
//
// f: 请求处理接口 protocol.RequestInterface 的实现
func (a *ActionMux) AddRequestInterface(f protocol.RequestInterface) error {
	request, err := a.checkRequestInterface(f)
	if err != nil {
		return err
	}
	// action已存在
	if _, ok := a.actionHandleFunc.Load(request.Action); ok {
		util.Logger.Warning("request mux: request action already exist")
		return ErrorActionExist
	}
	a.actionHandleFunc.Store(request.Action, f)
	return nil
}

// AddRegexRequestInterface 添加一个正则匹配动作请求接口
func (a *ActionMux) AddRegexRequestInterface(actionRegex string, f protocol.RequestInterface) error {
	_, err := a.checkRequestInterface(f)
	if err != nil {
		return err
	}
	re, err := regexp.Compile(actionRegex)
	if err != nil {
		return err
	}
	if _, ok := a.regexHandleFunc.Load(actionRegex); ok {
		util.Logger.Warning("request mux: request action regex already exist")
		return ErrorActionExist
	}
	newMux := new(regexActionMux)
	newMux.Re = re
	newMux.F = f
	a.regexHandleFunc.Store(actionRegex, newMux)
	return nil
}

// DeleteRequestInterface 删除一个请求处理接口
func (a *ActionMux) DeleteRequestInterface(action string) {
	a.actionHandleFunc.Delete(action)
}

func (a *ActionMux) DeleteRegexRequestInterface(actionRegex string) {
	a.regexHandleFunc.Delete(actionRegex)
}

// GetRequestInterface 获取一个请求处理接口
func (a *ActionMux) GetRequestInterface(action string) (f protocol.RequestInterface) {
	util.Logger.Debug("mux action request: " + action)
	ff, ok := a.actionHandleFunc.Load(action)
	if !ok {
		a.regexHandleFunc.Range(
			func(key, value interface{}) bool {
				if value.(*regexActionMux).Re.MatchString(action) {
					ff = value.(*regexActionMux).F
					ok = true
					return false
				}
				return true
			})
		if !ok {
			util.Logger.Warning("request mux: request action not exist")
			return nil
		}
	}
	return ff.(protocol.RequestInterface)
}
