package protocol

import (
	"errors"
	"reflect"
)

const (
	// StatusOk 正常
	StatusOk = "ok"
	// StatusFailed 失败
	StatusFailed = "failed"
)

// Reference: https://12.onebot.dev/connect/data-protocol/action-response/#0-ok
const (
	// ResponseCodeOk 请求成功
	// 当动作请求有效、动作执行成功时，返回码应为 0。
	ResponseCodeOk = 0
)

// Reference: https://12.onebot.dev/connect/data-protocol/action-response/#1xxxx-request-error
const (
	// ResponseCodeBadRequest 无效的动作请求
	// 格式错误（包括实现不支持 MessagePack 的情况）、必要字段缺失或字段类型错误
	ResponseCodeBadRequest = 10001
	// ResponseCodeUnsupportedAction 不支持的动作请求
	// OneBot 实现没有实现该动作
	ResponseCodeUnsupportedAction = 10002
	// ResponseCodeBadParam 无效的动作请求参数
	// 参数缺失或参数类型错误
	ResponseCodeBadParam = 10003
)

// Reference: https://12.onebot.dev/connect/data-protocol/action-response/#2xxxx-handler-error
const (
	// ResponseCodeBadHandler 	动作处理器实现错误
	// 没有正确设置响应状态等
	ResponseCodeBadHandler = 20001
)

//TODO: 补充所有的错误码

var (
	// ErrorInvalidResponse 该结构体不是一个OneBot动作相应
	ErrorInvalidResponse = errors.New("arg not a onebot response struct")
	// ErrorInvalidResponseRetCode 返回码无效
	ErrorInvalidResponseRetCode = errors.New("invalid response retcode")
)

var (
	// ResponseType 事件类型
	ResponseType = reflect.TypeOf(&Response{})
)

// ResponseCheck 检查响应是否合法
func ResponseCheck(r *Request, e any) error {
	// 不是指针或者不是结构体
	if reflect.TypeOf(e).Kind() != reflect.Ptr || reflect.TypeOf(e).Elem().Kind() != reflect.Struct {
		return ErrorInvalidResponse
	}
	responseId := -1
	// 判断是否存在Response类型字段
	t := reflect.TypeOf(e).Elem()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == ResponseType {
			responseId = i
			break
		}
	}
	// 不存在Response类型字段
	if responseId == -1 {
		// 该结构体不是一个OneBot动作相应
		return ErrorInvalidResponse
	}
	// 存在Response类型字段
	if !reflect.ValueOf(e).Elem().Field(responseId).IsValid() || reflect.ValueOf(e).Elem().Field(responseId).IsNil() {
		return ErrorInvalidResponse
	}
	newResponse := reflect.ValueOf(e).Elem().Field(responseId).Interface().(*Response)
	if newResponse.Retcode < 0 {
		// 响应码类型无效
		return ErrorInvalidResponseRetCode
	}
	// 通过响应码判断响应状态
	if newResponse.Retcode == 0 {
		newResponse.Status = StatusOk
	} else {
		newResponse.Status = StatusFailed
	}
	// 原样返回echo
	newResponse.Echo = r.Echo
	newResponse.requestID = r.GetID()
	return nil
}
