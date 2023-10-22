package v12

import (
	"errors"
	"fmt"
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	// ErrorResponseIDNotFound 响应ID未找到
	ErrorResponseIDNotFound = errors.New("response id not found")
)

// OneBotV12HttpConfig HTTP 连接配置
type OneBotV12HttpConfig struct {
	// Host HTTP 服务器监听 IP
	Host string `yaml:"host" json:"host" default:""`
	// Port HTTP 服务器监听端口
	Port int `yaml:"port" json:"port" default:""`
	// AccessToken 访问令牌
	AccessToken string `yaml:"access_token" json:"access_token" default:""`
	// TimeOut 超时时间
	// 如果开发者在该时间内没有给出相应, 自动503
	TimeOut int `yaml:"time_out" json:"time_out" default:"5000"`
	// EventEnable 是否启用 get_latest_events 元动作
	EventEnable bool `yaml:"event_enable" json:"event_enable" default:"true"`
	// EventBufferSize 事件缓冲区大小，超过该大小将会丢弃最旧的事件，0 表示不限大小
	EventBufferSize int `yaml:"event_buffer_size" json:"event_buffer_size" default:"500"`
}

type OneBotV12ConnectHttp struct {
	receiveFunc func([]byte) (string, error)
	config      *OneBotV12HttpConfig
	// requestCallBack chan []byte
	requestCallBack sync.Map
}

func (c *OneBotV12ConnectHttp) Send(b []byte, t int, e string) error {
	if t == sendTypeResponse {
		var try = 0
		// 尝试等待, 达到设置的最大超时时间
		for {
			if r, ok := c.requestCallBack.Load(e); ok {
				r.(chan []byte) <- b
				return nil
			}
			if try > c.config.TimeOut/200 {
				util.Logger.Warning("onebot v12 http server send response timeout")
				return ErrorResponseIDNotFound
			}
			try++
			time.Sleep(time.Millisecond * 200)
		}
	} else {
		util.Logger.Debug("onebot v12 http server throw event: " + e)
	}
	return nil
}

func (c *OneBotV12ConnectHttp) SetCallBackFunc(f OneBotV12ConnectCallBackFunc) {
	c.receiveFunc = f
}

func NewOneBotV12ConnectHttp(_ protocol.OneBotConfig, c any) (OneBotV12Connect, error) {
	newHttpConnect := OneBotV12ConnectHttp{}
	config := c.(OneBotV12HttpConfig)
	if config.TimeOut <= 0 {
		config.TimeOut = 5000
	}
	if config.EventBufferSize <= 0 {
		config.EventBufferSize = 65536
	}
	newHttpConnect.config = &config
	err := newHttpConnect.Start()
	if err != nil {
		return nil, err
	}
	return &newHttpConnect, nil
}

func (c *OneBotV12ConnectHttp) receiveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "github.com/FishZe/go-libonebot")
	w.Header().Set("OneBot-Version", "12")
	ticker := util.NewTicker()
	defer func() {
		if err := r.Body.Close(); err != nil {
			util.Logger.Warning("onebot v12 http server close body error: " + err.Error())
		}
		ticker.Stop()
		util.Logger.Debug("onebot v12 http server handle action cost: " + ticker.GetDurationString())
	}()
	util.Logger.Debug("onebot v12 http server receive request: " + r.Method + " " + r.Header.Get("Content-Type"))
	// 如果收到非 POST 请求，可以返回 HTTP 状态码 405 Method Not Allowed
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// 如果收到请求非 / 的路径，可以返回 HTTP 状态码 404 Not Found
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	/*
		标准定义两种 Content-Type 请求头：

		application/json：在请求体中使用 JSON 和 UTF-8 编码的字符串表示动作请求
		application/msgpack：在请求体中使用 MessagePack 编码的字节序列表示动作请求
		其中，application/json 是任何 OneBot 实现必须支持的，application/msgpack 是可选的。

		当收到上述任一种请求后，如果支持，应在响应头中设置相同的 Content-Type，并以相同的格式和编码返回动作响应。
	*/
	switch r.Header.Get("Content-Type") {
	case "application/json":
		w.Header().Set("Content-Type", "application/json")
		/*
			如果配置了 access_token 且不为空字符串，OneBot 实现必须：

			首先检查请求头中是否存在 Authorization 头，若存在，则判断其值是否等于 Bearer <access_token>（<access_token> 不需要对两边的空白字符进行裁剪），若等于则鉴权成功，否则鉴权失败；
			若不存在 Authorization 头，则继续检查是否存在 access_token URL query 参数，若存在，则判断其值是否等于 <access_token>，若等于则鉴权成功，否则鉴权失败；
			若以上均不存在，则鉴权失败。
		*/
		var authValue string
		if r.Header.Get("Authorization") == "" {
			authValue = r.URL.Query().Get("access_token")
		} else {
			authValue = r.Header.Get("Authorization")
		}
		if c.config.AccessToken != "" && authValue != fmt.Sprintf("Bearer<%s>", c.config.AccessToken) {
			// 如果鉴权失败，必须返回 HTTP 状态码 401 Unauthorized
			util.Logger.Debug("onebot v12 http server authorized failed")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		util.Logger.Debug("onebot v12 http server authorized successful")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			util.Logger.Warning("onebot v12 http server read body error: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		rId, err := c.receiveFunc(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ch := make(chan []byte, 0)
		c.requestCallBack.Store(rId, ch)
		defer func() {
			// 关闭channel, 删除映射关系
			util.Logger.Debug("onebot v12 http server delete response: " + rId)
			close(ch)
			c.requestCallBack.Delete(rId)
		}()
		util.Logger.Debug("onebot v12 http server waiting response: " + rId)
		// 一旦开始读取 HTTP 请求体，此后所有出错情形应通过动作响应的 retcode 字段区分，HTTP 状态码返回 200 OK
		w.WriteHeader(http.StatusOK)
		// 定时, 达到超时时间后返回
		tick := time.NewTicker(time.Second * time.Duration(c.config.TimeOut))
		// 等待response
		for {
			select {
			case b := <-ch:
				util.Logger.Debug("onebot v12 http server receive response: " + rId)
				_, err = io.WriteString(w, string(b))
				if err != nil {
					util.Logger.Error("onebot v12 http server respond error: " + err.Error())
				}
				return
			case <-tick.C:
				util.Logger.Debug("onebot v12 http server receive response timeout: " + rId)
				return
			}
		}
	case "application/msgpack":
		w.Header().Set("Content-Type", "application/msgpack")
	//TODO: 完善 msgpack 的支持
	default:
		// 如果收到不支持的 Content-Type 请求头，必须返回 HTTP 状态码 415 Unsupported Media Type
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

func (c *OneBotV12ConnectHttp) Start() error {
	//TODO: 检查端口是否被占用 检查ip和port是否合法
	if !util.CheckPortAvailable(c.config.Port) {
		return ErrorPortInUsed
	}
	/*
		OneBot 实现应该根据用户配置启动 HTTP 服务器，监听指定的 <host>:<port>，接受路径为 / 的 POST 请求，将 HTTP 请求体的内容解析为动作请求，处理后在 HTTP 响应体中返回动作响应。
	*/
	util.Logger.Debug("starting onebot v12 http server...")
	s := http.NewServeMux()
	s.HandleFunc("/", c.receiveHandler)
	mux := http.TimeoutHandler(s, time.Duration(c.config.TimeOut)*time.Millisecond, "")
	go func() {
		for {
			err := http.ListenAndServe(c.config.Host+":"+strconv.Itoa(c.config.Port), mux)
			util.Logger.Warning("onebot v12 http server error: " + err.Error())
			if err != nil {
				time.Sleep(time.Second * 5)
			}
		}

	}()
	return nil
}
