package v12

import (
	"fmt"
	"github.com/FishZe/go-libonebot/util"
	"github.com/fasthttp/websocket"
	"net/http"
	"time"
)

// OneBotV12WebsocketReverseConfig HTTP 连接配置
type OneBotV12WebsocketReverseConfig struct {
	// Url 反向 WebSocket 连接地址
	Url string
	// AccessToken 访问令牌
	AccessToken string
	// ReconnectInterval 反向 WebSocket 重连间隔，单位：毫秒，必须大于 0
	ReconnectInterval int
	// UserAgent HTTP 请求头中的 User-Agent 字段
	UserAgent string
	impl      string
}

// OneBotV12ConnectWebsocketReverse 反向 WebSocket 连接
type OneBotV12ConnectWebsocketReverse struct {
	receiveFunc func([]byte) (string, error)
	config      *OneBotV12WebsocketReverseConfig
	conn        *websocket.Conn
	sendChan    chan []byte
	done        chan struct{}
}

// Send 直接发送数据
func (c *OneBotV12ConnectWebsocketReverse) Send(b []byte, t int, e string) error {
	c.sendChan <- b
	return nil
}

// BindReceiveFunc 绑定接收函数
func (c *OneBotV12ConnectWebsocketReverse) BindReceiveFunc(f func([]byte) (string, error)) {
	c.receiveFunc = f
}

// startWebsocketClient 启动反向 WebSocket 客户端
func (c *OneBotV12ConnectWebsocketReverse) startWebsocketClient() {
	var err error
	for {
		// 定义header
		header := http.Header{}
		header.Set("User-Agent", c.config.UserAgent)
		header.Set("Sec-WebSocket-Protocol", "12."+c.config.impl)
		if c.config.AccessToken != "" {
			// AccessToken 不为空时，添加 Authorization 字段
			header.Set("Authorization", fmt.Sprintf("Bearer <%s>", c.config.AccessToken))
		}
		c.conn, _, err = websocket.DefaultDialer.Dial(c.config.Url, header)
		if err == nil {
			go func() {
				for {
					select {
					case msg := <-c.sendChan:
						err = c.conn.WriteMessage(websocket.TextMessage, msg)
						if err != nil {
							util.Logger.Warning("websocket send error: " + err.Error())
						}
					case <-c.done:
						return
					}
				}
			}()
			for {
				_, message, err := c.conn.ReadMessage()
				if err != nil {
					c.done <- struct{}{}
					break
				}
				if c.receiveFunc != nil {
					_, err := c.receiveFunc(message)
					if err != nil {
						util.Logger.Warning("websocket receive error: " + err.Error())
					}
				}
			}
		} else {
			util.Logger.Error("websocket connect error: " + err.Error())
		}
		// 等待重连
		time.Sleep(time.Duration(c.config.ReconnectInterval) * time.Millisecond)
	}
}

// Start 启动反向 WebSocket 连接
func (c *OneBotV12ConnectWebsocketReverse) Start() error {
	go c.startWebsocketClient()
	return nil
}

// NewOneBotV12ConnectWebsocketReverse 创建反向 WebSocket 连接
func NewOneBotV12ConnectWebsocketReverse(config *OneBotV12WebsocketReverseConfig) (*OneBotV12ConnectWebsocketReverse, error) {
	onebot := OneBotV12ConnectWebsocketReverse{
		config:   config,
		sendChan: make(chan []byte, 65535),
		done:     make(chan struct{}),
	}
	err := onebot.Start()
	if err != nil {
		return nil, err
	}
	return &onebot, nil
}
