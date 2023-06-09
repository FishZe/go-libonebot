package protocol

import (
	"errors"
	"github.com/FishZe/go-libonebot/util"
	"reflect"
)

const (
	// ActionTypeMeta 元事件
	ActionTypeMeta = "meta"
	// ActionTypeMessage 消息事件
	ActionTypeMessage = "message"
	// ActionTypeNotice 通知事件
	ActionTypeNotice = "notice"
	// ActionTypeRequest 请求事件
	ActionTypeRequest = "request"
)

var (
	// ErrorInvalidEvent 该结构体不是一个OneBot事件
	ErrorInvalidEvent = errors.New("arg not a onebot event strcut")
	// ErrorInValidEventType 事件类型无效
	ErrorInValidEventType = errors.New("event type is invalid")
)

var (
	// EventType 事件类型
	EventType = reflect.TypeOf(&Event{})
)

// EventCheck 检查事件是否合法
//
// s: Self 自身标识
// e: Event 事件 需要满足条件
func EventCheck(s Self, e any) (string, error) {
	// 寻找是否存在Event类型字段
	if reflect.TypeOf(e).Kind() != reflect.Ptr || reflect.TypeOf(e).Elem().Kind() != reflect.Struct {
		util.Logger.Warning("EventCheck: arg not a onebot event struct")
		return "", ErrorInvalidEvent
	}
	eventId := -1
	t := reflect.TypeOf(e).Elem()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == EventType {
			eventId = i
			break
		}
	}
	// 不存在Event类型字段
	if eventId == -1 {
		// 该结构体不是一个OneBot事件
		util.Logger.Warning("EventCheck: arg not a onebot event struct")
		return "", ErrorInvalidEvent
	}
	// 存在Event类型字段
	if !reflect.ValueOf(e).Elem().Field(eventId).IsValid() || reflect.ValueOf(e).Elem().Field(eventId).IsNil() {
		util.Logger.Warning("EventCheck: event not valid")
		return "", ErrorInvalidEvent
	}
	newEvent := reflect.ValueOf(e).Elem().Field(eventId).Interface().(*Event)
	if newEvent.Type == "" || (newEvent.Type != ActionTypeMessage && newEvent.Type != ActionTypeMeta && newEvent.Type != ActionTypeNotice && newEvent.Type != ActionTypeRequest) {
		// 事件类型无效
		util.Logger.Warning("EventCheck: event not valid")
		return "", ErrorInValidEventType
	}
	// 事件ID和时间戳为空时, 设置默认值
	if newEvent.ID == "" {
		newEvent.ID = util.GetUUID()
	}
	if newEvent.Time <= 0 {
		newEvent.Time = util.GetTimeStampFloat64()
	}
	newEvent.Self = s
	util.Logger.Debug("event check: " + newEvent.ID + " success")
	return newEvent.ID, nil
}
