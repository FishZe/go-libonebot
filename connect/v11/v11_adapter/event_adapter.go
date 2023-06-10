package v11_adapter

import (
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
	"reflect"
	"strconv"
)

const (
	V11PostTypeMessage   = "message"
	V11PostTypeNotice    = "notice"
	V11PostTypeRequest   = "request"
	V11PostTypeMetaEvent = "meta_event"
)

type V11EventHead struct {
	// Time 时间戳
	Time int64 `json:"time"`
	// SelfId 机器人ID
	SelfId int64 `json:"self_id"`
	// PostType 类型
	PostType string `json:"post_type"`
	// MessageType 消息类型
	MessageType string `json:"message_type"`
	// MetaEventType 元事件类型
	MetaEventType string `json:"meta_event_type"`
	//  SubType 事件子类型
	SubType string `json:"sub_type"`
}

type V11Event struct {
	*V11EventHead
	// Data 数据
	Data map[string]interface{}
}

func EventHead12To11(event *protocol.Event) (res *V11EventHead) {
	res.Time = int64(event.Time)
	if event.Type == protocol.EventTypeMeta {
		util.Logger.Debug("V11 adapter: meta -> meta_event")
		res.PostType = V11PostTypeMetaEvent
	} else {
		res.PostType = event.Type
	}
	id, err := strconv.ParseInt(event.Self.UserId, 10, 64)
	if err != nil {
		id = 0
	}
	res.SelfId = id
	return
}

// Event12To11
// TODO: 未完成
func Event12To11(event any) *V11Event {
	res := V11Event{
		Data: make(map[string]interface{}),
	}
	//
	if reflect.TypeOf(event).Kind() != reflect.Ptr || reflect.TypeOf(event).Elem().Kind() != reflect.Struct {
		util.Logger.Warning("V11 adapter: arg not a V12 onebot event struct")
		return nil
	}
	eventId := -1
	t := reflect.TypeOf(event).Elem()
	v := reflect.ValueOf(event).Elem()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == protocol.EventType {
			eventId = i
			break
		}
	}
	if eventId == -1 {
		util.Logger.Warning("V11 adapter: arg not a V12 onebot event struct")
		return nil
	}
	switch v.Type() {
	case reflect.TypeOf(&protocol.MetaEventConnect{}):
	case reflect.TypeOf(&protocol.MetaEventHeartbeat{}):
	case reflect.TypeOf(&protocol.MetaEventStatusUpdate{}):
	case reflect.TypeOf(&protocol.MessageEventPrivate{}):
	case reflect.TypeOf(&protocol.MessageEventGroup{}):
	case reflect.TypeOf(&protocol.MessageEventChannel{}):
	case reflect.TypeOf(&protocol.NoticeEventChannelCreate{}):

	}

	// 先直接塞进来
	event12 := v.Field(eventId).Interface().(*protocol.Event)
	res.V11EventHead = EventHead12To11(event12)
	for i := 0; i < t.NumField(); i++ {
		if i == eventId {
			continue
		}

	}
	return &res
}
