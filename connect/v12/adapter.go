package v12

import (
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
	"reflect"
)

// OneBotV12Response V12响应
type OneBotV12Response struct {
	*protocol.Response
	Data any `json:"data"`
}

// ConnectSendEvent 发送事件
func (o *OneBotV12) ConnectSendEvent(e any, eid string) error {
	s, err := util.Json.Marshal(e)
	if err != nil {
		util.Logger.Warning("onebot v12 marshal event error: " + err.Error())
		return err
	}
	util.Logger.Debug("onebot v12 send event: " + eid)
	return o.connect.Send(s, SendTypeEvent, eid)
}

// ConnectSendResponse 发送响应
func (o *OneBotV12) ConnectSendResponse(e any) error {
	// 不是指针或者不是结构体
	if reflect.TypeOf(e).Kind() != reflect.Ptr || reflect.TypeOf(e).Elem().Kind() != reflect.Struct {
		return protocol.ErrorInvalidResponse
	}
	// 创建V12 Response
	newV12Response := OneBotV12Response{
		Data: make(map[string]any),
	}
	findSlice := false
	t := reflect.TypeOf(e).Elem()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == protocol.ResponseType {
			// Response字段
			if !reflect.ValueOf(e).Elem().Field(i).IsValid() || reflect.ValueOf(e).Elem().Field(i).IsNil() {
				// 这个Response字段不合法
				return protocol.ErrorInvalidResponse
			}
			newV12Response.Response = reflect.ValueOf(e).Elem().Field(i).Interface().(*protocol.Response)
		} else if t.Field(i).Tag.Get("json") != "" {
			// 其他字段, 判断是否存在json tag 不存在则跳过
			newV12Response.Data.(map[string]any)[t.Field(i).Tag.Get("json")] = reflect.ValueOf(e).Elem().Field(i).Interface()
		} else if reflect.ValueOf(e).Elem().Field(i).Kind() == reflect.Slice {
			// 判断是否存在Slice
			findSlice = true
		}
	}
	if len(newV12Response.Data.(map[string]any)) == 0 && findSlice {
		// 说明data字段为切片
		newV12Response.Data = []any{}
		for i := 0; i < t.NumField(); i++ {
			if reflect.ValueOf(e).Elem().Field(i).Kind() == reflect.Slice {
				newV12Response.Data = reflect.ValueOf(e).Elem().Field(i).Interface()
			}
		}
	}
	s, err := util.Json.Marshal(newV12Response)
	if err != nil {
		util.Logger.Warning("onebot v12 marshal response error: " + err.Error())
		return err
	}
	util.Logger.Debug("onebot v12 send response: " + newV12Response.GetID())
	return o.connect.Send(s, sendTypeResponse, newV12Response.GetID())
}

// receiveMessage 接收消息
func (o *OneBotV12) receiveMessage(b []byte) (string, error) {
	x := protocol.RawRequestType{}
	err := util.Json.Unmarshal(b, &x)
	if err != nil {
		util.Logger.Warning("onebot v12 decode message error: " + err.Error())
		return "", err
	}
	x.Request.NewID()
	x.Request.ConnectionUUID = o.GetUUID()
	if len(o.botList) == 1 {
		// 试图补充self字段
		// 方便后续使用
		o.botId.Range(func(k any, v any) bool {
			x.Self = k.(protocol.Self)
			return false
		})
		// 只连接了一个, self字段不是必须的
		if ch, ok := o.botRequestChan.Load(0); ok {
			ch.(chan protocol.RawRequestType) <- x
		}
	} else {
		// 多个, self字段必须存在
		if c, ok := o.botId.Load(x.Self); ok {
			if ch, ok := o.botRequestChan.Load(c); ok {
				ch.(chan protocol.RawRequestType) <- x
			} else {
				util.Logger.Warning("onebot v12 receive message error: chan not exist")
				go func(id string) {
					e := protocol.NewEmptyResponse(protocol.ResponseCodeInternalHandlerError)
					e.SetID(x.Request.GetID())
					_ = o.ConnectSendResponse(e)
				}(x.Request.GetID())
			}
		} else {
			util.Logger.Warning("onebot v12 receive message error: self not exist")
			go func(id string) {
				e := protocol.NewEmptyResponse(protocol.ResponseCodeUnknownSelf)
				e.SetID(id)
				_ = o.ConnectSendResponse(e)
			}(x.Request.GetID())
		}
	}
	util.Logger.Debug("onebot v12 receive message: " + x.Request.GetID())
	return x.Request.GetID(), nil
}
