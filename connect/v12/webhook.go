package v12

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// OneBotV12WebHookConfig HTTP 连接配置
type OneBotV12WebHookConfig struct {
	// Url Webhook 上报地址
	Url string
	//AccessToken 访问令牌
	AccessToken string
	// TimeOut 上报请求超时时间，单位：毫秒，0 表示不超时
	TimeOut int
	// UserAgent
	UserAgent string
	impl      string
}

// OneBotV12ConnectWebhook OneBot v12 webhook 连接
// 这个！我没测试！
// TODO: 测试
type OneBotV12ConnectWebhook struct {
	receiveFunc func([]byte) (string, error)
	config      *OneBotV12WebHookConfig
}

// returnRequest 返回请求 分割
func (c *OneBotV12ConnectWebhook) returnRequest(b []byte) error {
	// 响应体的内容解析为动作请求列表
	var list []interface{}
	err := json.Unmarshal(b, &list)
	if err != nil {
		return err
	}
	for _, v := range list {
		// 依次处理
		r, err := json.Marshal(v)
		if err == nil {
			_, _ = c.receiveFunc(r)
		}
	}
	return nil
}

// Send 发送事件
func (c *OneBotV12ConnectWebhook) Send(b []byte, t int, e string) error {
	// OneBot 实现应该根据用户配置，在发生事件时，向指定的 url 使用 POST 请求推送事件，并根据情况将 HTTP 响应体的内容解析为动作请求列表，依次处理，丢弃动作响应。
	if t == sendTypeResponse {
		return nil
	} else if t == SendTypeEvent {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: time.Duration(c.config.TimeOut) * time.Millisecond,
		}
		req, err := http.NewRequest("POST", c.config.Url, bytes.NewBuffer(b))
		req.Header = http.Header{
			// TODO: 支持 application/msgpack
			"Content-Type":     {"application/json"},
			"User-Agent":       {c.config.UserAgent},
			"X-OneBot-Version": {"12"},
			"X-Impl":           {c.config.impl},
		}
		if c.config.AccessToken != "" {
			req.Header.Add("Authorization", fmt.Sprintf("Bearer <%s>", c.config.AccessToken))
		}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
		if resp.StatusCode != 200 && resp.StatusCode != 204 {
			return errors.New("onebot v12 webhook status code error: " + strconv.Itoa(resp.StatusCode) + " " + resp.Status)
		}
		s, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return c.returnRequest(s)
	}
	return nil
}

// BindReceiveFunc 绑定接收函数
func (c *OneBotV12ConnectWebhook) BindReceiveFunc(f func([]byte) (string, error)) {
	c.receiveFunc = f
}

// NewOneBotV12ConnectWebhook 创建一个 OneBot v12 webhook 连接
func NewOneBotV12ConnectWebhook(config *OneBotV12WebHookConfig) (*OneBotV12ConnectWebhook, error) {
	onebot := OneBotV12ConnectWebhook{config: config}
	return &onebot, nil
}
