package protocol

import (
	"errors"
	"github.com/FishZe/go-libonebot/util"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

var (
	// ErrorInvalidRequest 该结构体不是一个OneBot事件
	ErrorInvalidRequest = errors.New("arg not a onebot request struct")
)

var (
	// RequestType 事件类型
	RequestType = reflect.TypeOf(&Request{})
)

// Request 动作请求
type Request struct {
	// Action 动作名称
	Action string `json:"action"`
	// Echo 原样返回
	Echo string `json:"echo"`
	// Self 事件发起Bot
	Self Self `json:"self"`
	// ConnectionUUID 内部使用: 判断发出请求的连接UUID 用于返回时寻找连接
	ConnectionUUID string
	//RequestID 内部标识字段
	requestID string
}

func (r *Request) NewID() {
	r.requestID = util.GetUUID()
}

func (r *Request) GetID() string {
	return r.requestID
}

// RawRequestType 最原始的动作请求, 用于兼容和各种连接交互
type RawRequestType struct {
	// *Request 请求头
	*Request
	// Param 各种参数
	Param map[string]any `json:"params"`
}

// RawRequestTypeToRequest 将RawRequestType转换为Request
func RawRequestTypeToRequest(e any, r RawRequestType) error {
	if reflect.TypeOf(e).Kind() != reflect.Ptr || reflect.TypeOf(e).Elem().Kind() != reflect.Struct {
		return ErrorInvalidRequest
	}
	t := reflect.TypeOf(e).Elem()
	v := reflect.ValueOf(e).Elem()
	// 类型不对
	if reflect.ValueOf(e).Elem().Kind() != reflect.Struct || reflect.ValueOf(e).Elem() == reflect.Zero(t) {
		return ErrorInvalidRequest
	}
	findRequest := false
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == RequestType {
			// 找到Request类型字段
			if !v.Field(i).IsValid() {
				// 不符合
				return ErrorInvalidRequest
			} else if v.Field(i).IsNil() {
				// 未设置
				v.Field(i).Set(reflect.ValueOf(&Request{}))
			}
			findRequest = true
			// 修改Request
			v.Field(i).Interface().(*Request).Action = r.Action
			v.Field(i).Interface().(*Request).Echo = r.Echo
			v.Field(i).Interface().(*Request).Self = r.Self
			v.Field(i).Interface().(*Request).ConnectionUUID = r.ConnectionUUID
		}
	}
	// 未找到Request类型字段
	if !findRequest {
		// 该结构体不是一个OneBot动作请求
		return ErrorInvalidResponse
	}
	// 赋值
	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: false,
		Result:      e,
	}); err == nil {
		if err = decoder.Decode(r.Param); err == nil {
			return nil
		}
		return err
	} else {
		return err
	}
}

// RequestInterface 动作请求接口
type RequestInterface interface {
	// Do 执行动作
	// e 返回的Response 需要符合Response的类型
	Do() (e any)
	// New 返回一个新的RequestInterface
	New() any
}
