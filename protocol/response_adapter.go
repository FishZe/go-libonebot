package protocol

import (
	"errors"
	"github.com/FishZe/go-libonebot/util"
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
	// ResponseCodeUnsupportedParam 不支持的动作请求参数
	// OneBot 实现没有实现该参数的语义
	ResponseCodeUnsupportedParam = 10004
	// ResponseCodeUnsupportedSegment 不支持的消息段类型
	//  OneBot 实现没有实现该消息段类型
	ResponseCodeUnsupportedSegment = 10005
	// ResponseCodeBadSegmentData 无效的消息段参数
	// 参数缺失或参数类型错误
	ResponseCodeBadSegmentData = 10006
	// ResponseCodeUnsupportedSegmentData 不支持的消息段参数
	// OneBot 实现没有实现该参数的语义
	ResponseCodeUnsupportedSegmentData = 10007
	// ResponseCodeWhoAmI 未指定机器人账号
	// OneBot 实现在单个 OneBot Connect 连接上支持多个机器人账号，但动作请求未指定要使用的账号
	ResponseCodeWhoAmI = 10101
	// ResponseCodeUnknownSelf 未知的机器人账号
	// 动作请求指定的机器人账号不存在
	ResponseCodeUnknownSelf = 10102
)

// Reference: https://12.onebot.dev/connect/data-protocol/action-response/#2xxxx-handler-error
const (
	// ResponseCodeBadHandler 	动作处理器实现错误
	// 没有正确设置响应状态等
	ResponseCodeBadHandler = 20001
	// ResponseCodeInternalHandlerError 动作处理器运行时抛出异常
	// OneBot 实现内部发生了未捕获的意料之外的异常
	ResponseCodeInternalHandlerError = 20002
)

// Reference: https://12.onebot.dev/connect/data-protocol/action-response/#3xxxx-execution-error
/*
	3xxxx 动作执行错误（Execution Error）
	当动作请求有效，但动作执行失败时，返回码建议为下表中的一种，其中低三位可以由 OneBot 实现自行定义：

	错误码	 错误名	            说明	          备注
	31xxx	 Database Error	    数据库错误	  如数据库查询失败等
	32xxx	 Filesystem Error	文件系统错误	  如读取或写入文件失败等
	33xxx	 Network Error	    网络错误	      如下载文件失败等
	34xxx	 Platform Error	    机器人平台错误	  如由于机器人平台限制导致消息发送失败等
	35xxx	 Logic Error	    动作逻辑错误	  如尝试向不存在的用户发送消息等
	36xxx	 I Am Tired	        我不想干了	  一位 OneBot 实现决定罢工

*/

// Reference: https://12.onebot.dev/connect/data-protocol/action-response/#4xxxx5xxxx
/*

	4xxxx、5xxxx 保留错误段
	这两段返回码为保留段，OneBot 实现不应该使用。

*/

// Reference: https://12.onebot.dev/connect/data-protocol/action-response/#6xxxx9xxxx
/*

	对于 3xxxx 无法涵盖的情形，OneBot 实现可以自由使用其它错误段来表示。

*/

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

// Response 动作响应
// Reference: https://12.onebot.dev/connect/data-protocol/action-response/
type Response struct {
	// Status 响应状态
	// 必须为 StatusOk 或 StatusFailed
	Status string `json:"status"`
	// RetCode 返回码 对应上文的 ResponseCodeXXX
	Retcode int `json:"retcode"`
	// Message 错误信息，当动作执行失败时，建议在此填写人类可读的错误信息，当执行成功时，应为空字符串
	Message string `json:"message"`
	// Echo 应原样返回动作请求中的 echo 字段值
	Echo string `json:"echo"`
	// requestID 内部标识字段
	requestID string
}

// GetID 获取请求ID
func (r *Response) GetID() string {
	return r.requestID
}

// SetID 设置请求ID
func (r *Response) SetID(s string) {
	r.requestID = s
}

// ResponseCheck 检查响应是否合法
func ResponseCheck(r *Request, e any) error {
	// 不是指针或者不是结构体
	if reflect.TypeOf(e).Kind() != reflect.Ptr || reflect.TypeOf(e).Elem().Kind() != reflect.Struct {
		util.Logger.Warning("Response Check: arg not a onebot response struct")
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
		util.Logger.Warning("Response Check: arg not a onebot response struct")
		return ErrorInvalidResponse
	}
	// 存在Response类型字段
	if !reflect.ValueOf(e).Elem().Field(responseId).IsValid() || reflect.ValueOf(e).Elem().Field(responseId).IsNil() {
		// 该结构体不是一个OneBot动作相应
		util.Logger.Warning("Response Check: response not valid")
		return ErrorInvalidResponse
	}
	newResponse := reflect.ValueOf(e).Elem().Field(responseId).Interface().(*Response)
	if newResponse.Retcode < 0 {
		// 响应码类型无效
		util.Logger.Warning("Response Check: response code not valid")
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
	util.Logger.Debug("response " + newResponse.requestID + " check success")
	return nil
}
