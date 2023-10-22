package v12

import (
	"fmt"
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
	"github.com/fasthttp/websocket"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// OneBotV12WebsocketConfig HTTP 连接配置
type OneBotV12WebsocketConfig struct {
	// Host WebSocket 服务器监听 IP
	Host string `yaml:"host" json:"host" default:""`
	// Port WebSocket 服务器监听端口
	Port int `yaml:"port" json:"port" default:""`
	// AccessToken WebSocket 服务器监听端口
	AccessToken string `yaml:"access_token" json:"access_token" default:""`
	// TimeOut 超时时间
	TimeOut int `yaml:"time_out" json:"time_out" default:"5000"`
}

// OneBotV12ConnectWebsocket websocket链接
type OneBotV12ConnectWebsocket struct {
	config      OneBotV12WebsocketConfig
	receiveFunc func([]byte) (string, error)
	connections sync.Map
	waitResp    sync.Map
}

// websocketConnect 每一个websocket连接
type websocketConnect struct {
	sendChan chan []byte
	done     chan struct{}
	tickers  sync.Map
}

// Send 发送消息
func (c *OneBotV12ConnectWebsocket) Send(b []byte, t int, e string) error {
	if t == sendTypeResponse {
		// requestID -> connectionID
		if r, ok := c.waitResp.Load(e); ok {
			// connectionID -> connection
			if con, ok := c.connections.Load(r); ok {
				nowConn := con.(*websocketConnect)
				// requestID -> *util.Ticker
				if t, ok := nowConn.tickers.Load(e); ok {
					// 查找是否存在计时器
					nowConn.sendChan <- b
					tic := t.(*util.Ticker)
					tic.Stop()
					util.Logger.Debug("onebot v12 websocket send :" + e + " use " + tic.GetDurationString())
					nowConn.tickers.Delete(e)
				} else {
					// 动作超时
					util.Logger.Warning("onebot v12 websocket send :" + e + " timeout")
				}
			}
		}
	} else {
		c.connections.Range(func(key, value interface{}) bool {
			value.(*websocketConnect).sendChan <- b
			return true
		})
	}
	return nil
}

// SetCallBackFunc 绑定接收函数
func (c *OneBotV12ConnectWebsocket) SetCallBackFunc(f OneBotV12ConnectCallBackFunc) {
	c.receiveFunc = f
}

// NewOneBotV12ConnectWebsocket 创建websocket连接实现
func NewOneBotV12ConnectWebsocket(_ protocol.OneBotConfig, c any) (OneBotV12Connect, error) {
	onebot := OneBotV12ConnectWebsocket{}
	config := c.(OneBotV12WebsocketConfig)
	if config.TimeOut <= 0 {
		config.TimeOut = 5000
	}
	onebot.config = config
	err := onebot.Start()
	if err != nil {
		return nil, err
	}
	return &onebot, nil
}

// receiveHandler websocket连接处理函数
func (c *OneBotV12ConnectWebsocket) receiveHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		// TODO:
	}()
	// 如果收到请求非 / 的路径，可以返回 HTTP 状态码 404 Not Found
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var authValue string
	if r.Header.Get("Authorization") == "" {
		authValue = r.URL.Query().Get("access_token")
	} else {
		authValue = r.Header.Get("Authorization")
	}
	if c.config.AccessToken != "" && authValue != fmt.Sprintf("Bearer<%s>", c.config.AccessToken) {
		// 如果鉴权失败，必须返回 HTTP 状态码 401 Unauthorized
		util.Logger.Debug("onebot v12 websocket server authorized failed")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	util.Logger.Debug("onebot v12 websocket server authorized successful")
	// 鉴权成功 upgrade
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// upgrade
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		util.Logger.Error("onebot v12 websocket server upgrade failed: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	util.Logger.Debug("onebot v12 websocket server connect successful")
	// 新建链接
	newConn := websocketConnect{
		sendChan: make(chan []byte, 65535),
		done:     make(chan struct{}),
	}
	nowWaitID := make([]string, 0)
	nowID := util.GetUUID()
	c.connections.Store(nowID, &newConn)
	util.Logger.Debug("onebot v12 websocket server new connection: " + nowID)
	defer func() {
		// 关闭连接
		util.Logger.Debug("onebot v12 websocket server close connection: " + nowID)
		err = conn.Close()
		if err != nil {
			util.Logger.Error("onebot v12 websocket server close failed: " + err.Error())
		}
		// 清除连接
		c.connections.Delete(nowID)
		for _, v := range nowWaitID {
			c.waitResp.Delete(v)
			if _, ok := newConn.tickers.Load(v); ok {
				newConn.tickers.Delete(v)
			}
		}
		// 清空列表 gc
		nowWaitID = nil
		// 关闭通道
		close(newConn.sendChan)
		close(newConn.done)
	}()
	go func() {
		for {
			select {
			case <-newConn.done:
				// 退出信号
				util.Logger.Debug("onebot v12 websocket server done")
				return
			case msg := <-newConn.sendChan:
				// 发送消息
				util.Logger.Debug("onebot v12 websocket server send msg")
				err = conn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					util.Logger.Warning("onebot v12 websocket server send failed: " + err.Error())
				}
			}

		}
	}()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			util.Logger.Error("onebot v12 websocket server read failed: " + err.Error())
			// 发送关闭信号
			newConn.done <- struct{}{}
			return
		}
		id, err := c.receiveFunc(message)
		if err != nil {
			util.Logger.Error("onebot v12 websocket server receive failed: " + err.Error())
		}
		util.Logger.Debug("onebot v12 websocket server wait response: " + id)
		// 存储requestID -> connectionID
		c.waitResp.Store(id, nowID)
		// 加入列表
		nowWaitID = append(nowWaitID, id)
		// 储存计时器
		newConn.tickers.Store(id, util.NewTicker())
		go func() {
			time.Sleep(time.Millisecond * time.Duration(c.config.TimeOut))
			if _, ok := newConn.tickers.Load(id); ok {
				// 超时
				util.Logger.Warning("onebot v12 websocket server timeout: " + id)
				// 删除计时器
				newConn.tickers.Delete(id)
				// 删除列表
				for i, v := range nowWaitID {
					if v == id {
						nowWaitID = append(nowWaitID[:i], nowWaitID[i+1:]...)
						break
					}
				}
				// 删除存储
				c.waitResp.Delete(id)
			}
		}()
	}
}

// Start 启动websocket服务
func (c *OneBotV12ConnectWebsocket) Start() error {
	//TODO: 检查端口是否被占用 检查ip和port是否合法
	if !util.CheckPortAvailable(c.config.Port) {
		return ErrorPortInUsed
	}
	util.Logger.Debug("starting onebot v12 websocket server...")
	s := http.NewServeMux()
	s.HandleFunc("/", c.receiveHandler)
	go func() {
		for {
			util.Logger.Debug(c.config.Host + ":" + strconv.Itoa(c.config.Port))
			err := http.ListenAndServe(c.config.Host+":"+strconv.Itoa(c.config.Port), s)
			util.Logger.Warning("onebot v12 websocket server error: " + err.Error())
			if err != nil {
				time.Sleep(time.Second * 5)
			}
		}

	}()
	return nil
}
