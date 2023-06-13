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
	// ErrorRequestIsNil 请求为空
	ErrorRequestIsNil = errors.New("request is nil")
	// ErrorRequestNotMatch 请求不匹配
	ErrorRequestNotMatch = errors.New("request not match")
	// ErrorActionEmpty 动作为空
	ErrorActionEmpty = errors.New("action is empty")
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

func (r *Request) SetID(s string) {
	r.requestID = s
}

// RawRequestType 最原始的动作请求, 用于兼容和各种连接交互
type RawRequestType struct {
	// *Request 请求头
	*Request
	// Param 各种参数
	Param map[string]any `json:"params"`
}

// RequestAdapter 将RawRequestType转换为Request
func RequestAdapter(e any, r RawRequestType) error {
	if e == nil || reflect.TypeOf(e).Kind() != reflect.Ptr || reflect.TypeOf(e).Elem().Kind() != reflect.Struct {
		util.Logger.Warning("RawRequestType To Request: arg not a onebot request struct")
		return ErrorInvalidRequest
	}
	t := reflect.TypeOf(e).Elem()
	v := reflect.ValueOf(e).Elem()
	findRequest := false
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == RequestType {
			// 找到Request类型字段
			if !v.Field(i).IsValid() || v.Field(i).IsNil() {
				// 不符合
				return ErrorRequestIsNil
			}
			findRequest = true
			if v.Field(i).Interface().(*Request).Action == "" || r.Action == "" {
				return ErrorActionEmpty
			}
			/*
				if v.Field(i).Interface().(*Request).Action != r.Action {
					return ErrorRequestNotMatch
				}*/
			// 修改Request
			/*
				v.Field(i).Interface().(*Request).Action = r.Action
				v.Field(i).Interface().(*Request).Echo = r.Echo
				v.Field(i).Interface().(*Request).Self = r.Self
				v.Field(i).Interface().(*Request).ConnectionUUID = r.ConnectionUUID
			*/
			newRequest := Request{
				Action:         r.Action,
				Echo:           r.Echo,
				Self:           r.Self,
				ConnectionUUID: r.ConnectionUUID,
				requestID:      r.requestID,
			}
			v.Field(i).Set(reflect.ValueOf(&newRequest))
		}
	}
	// 未找到Request类型字段
	if !findRequest {
		// 该结构体不是一个OneBot动作请求
		return ErrorInvalidRequest
	}
	// 赋值

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: false,
		Result:      e,
		TagName:     "json",
		ZeroFields:  true,
	}); err == nil {
		if err = decoder.Decode(r.Param); err == nil {
			util.Logger.Debug("request " + r.requestID + " decode & check success")
			return nil
		}
		util.Logger.Warning("RawRequestType To Request: decode error: " + err.Error())
		return err
	} else {
		util.Logger.Warning("RawRequestType To Request: decode error: " + err.Error())
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
