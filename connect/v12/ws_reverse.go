package v12

import (
	"fmt"
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
	"github.com/fasthttp/websocket"
	"net/http"
	"sync"
	"time"
)

// OneBotV12WebsocketReverseConfig HTTP 连接配置
type OneBotV12WebsocketReverseConfig struct {
	// Url 反向 WebSocket 连接地址
	Url string `yaml:"url" json:"url" default:""`
	// AccessToken 访问令牌
	AccessToken string `yaml:"access_token" json:"access_token" default:""`
	// ReconnectInterval 反向 WebSocket 重连间隔，单位：毫秒，必须大于 0
	ReconnectInterval int `yaml:"reconnect_interval" json:"reconnect_interval" default:"5000"`
	// UserAgent HTTP 请求头中的 User-Agent 字段
	UserAgent string `yaml:"user_agent" json:"user_agent" default:"go-libonebot"`
	// TimeOut
	TimeOut int `yaml:"time_out" json:"time_out" default:"5000"`
	// BufferSize 缓存事件和响应
	BufferSize int `yaml:"buffer_size" json:"buffer_size" default:"500"`
	impl       string
}

// OneBotV12ConnectWebsocketReverse 反向 WebSocket 连接
type OneBotV12ConnectWebsocketReverse struct {
	receiveFunc func([]byte) (string, error)
	config      *OneBotV12WebsocketReverseConfig
	conn        *websocket.Conn
	sendChan    chan []byte
	done        chan struct{}
	tickers     sync.Map
}

// Send 直接发送数据
func (c *OneBotV12ConnectWebsocketReverse) Send(b []byte, t int, e string) error {
	if t == sendTypeResponse {
		if v, ok := c.tickers.Load(e); ok {
			c.sendChan <- b
			ticker := v.(*util.Ticker)
			ticker.Stop()
			util.Logger.Debug("onebot v12 websocket reverse send :" + e + " use " + ticker.GetDurationString())
			c.tickers.Delete(e)
		} else {
			// 动作超时
			util.Logger.Warning("onebot v12 websocket reverse send :" + e + " timeout")
		}
	} else {
		c.sendChan <- b
		util.Logger.Debug("onebot v12 websocket reverse send :" + e)
	}

	return nil
}

// SetCallBackFunc 绑定接收函数
func (c *OneBotV12ConnectWebsocketReverse) SetCallBackFunc(f OneBotV12ConnectCallBackFunc) {
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
			util.Logger.Debug("onebot v12 websocket reverse connect success")
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
				// 收到消息
				if c.receiveFunc != nil {
					id, err := c.receiveFunc(message)
					if err != nil {
						util.Logger.Warning("websocket receive error: " + err.Error())
					}
					c.tickers.Store(id, util.NewTicker())
					// 记录动作id
					go func(id string) {
						time.Sleep(time.Duration(c.config.TimeOut) * time.Millisecond)
						if _, ok := c.tickers.Load(id); ok {
							c.tickers.Delete(id)
							// 超时
							util.Logger.Warning("onebot v12 websocket reverse receive timeout: " + id)
						}
					}(id)
				}
			}
		} else {
			util.Logger.Warning("websocket connect error: " + err.Error())
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
func NewOneBotV12ConnectWebsocketReverse(oc protocol.OneBotConfig, c any) (OneBotV12Connect, error) {
	config := c.(OneBotV12WebsocketReverseConfig)
	if config.ReconnectInterval <= 0 {
		// 默认5000ms
		config.ReconnectInterval = 5000
	}
	if config.UserAgent == "" {
		// 默认go-libonebot
		config.UserAgent = "github.com/FishZe/go-libonebot"
	}
	if config.TimeOut <= 0 {
		// 默认10000ms
		config.TimeOut = 10000
	}
	if config.BufferSize <= 0 {
		config.BufferSize = 65535
	}
	config.impl = oc.Implementation
	onebot := OneBotV12ConnectWebsocketReverse{
		config:   &config,
		sendChan: make(chan []byte, config.BufferSize),
		done:     make(chan struct{}),
	}
	err := onebot.Start()
	if err != nil {
		return nil, err
	}
	return &onebot, nil
}
