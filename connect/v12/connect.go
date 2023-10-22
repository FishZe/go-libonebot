package v12

import (
	"errors"
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
	"sync"
)

const (
	// ConnectTypeHttp http连接
	ConnectTypeHttp = -1
	// ConnectTypeWebhook webhook 连接
	ConnectTypeWebhook = -2
	// ConnectTypeWebSocket websocket连接
	ConnectTypeWebSocket = -3
	// ConnectTypeWebSocketReverse 反向websocket连接
	ConnectTypeWebSocketReverse = -4
)

const (
	sendTypeResponse = 1 << 0
	SendTypeEvent    = 1 << 1
)

type OneBotV12ConnectType int
type OneBotV12ConnectResisterFunc func(protocol.OneBotConfig, any) (OneBotV12Connect, error)
type OneBotV12ConnectCallBackFunc func([]byte) (string, error)

var (
	// ErrorConnectionNotFound 连接不存在
	ErrorConnectionNotFound = errors.New("connection not found")
	// ErrorConnectTypeExist 连接类型已经存在
	ErrorConnectTypeExist = errors.New("connect type exist")
	// ErrorConnectionNotV12 连接了错误的版本
	ErrorConnectionNotV12 = errors.New("the connection is not oneBot v12")
	// ErrorBotExist bot已经存在
	ErrorBotExist = errors.New("bot already exists")
	// ErrorBotNotExist bot不存在
	ErrorBotNotExist = errors.New("bot does not exist")
	// ErrorPortInUsed 端口被占用
	ErrorPortInUsed = errors.New("the port is in used")
)

var (
	// OneBotV12ConnectImplementations 连接实现 OneBotV12ConnectType -> OneBotV12ConnectResisterFunc
	OneBotV12ConnectImplementations sync.Map
)

// OneBotV12Connect 连接接口
type OneBotV12Connect interface {
	Send([]byte, int, string) error
	SetCallBackFunc(OneBotV12ConnectCallBackFunc)
}

// OneBotV12 OneBot V12连接
type OneBotV12 struct {
	// 连接类型
	connectType OneBotV12ConnectType
	// 连接的bot
	connectionUUID string
	// botList 机器人列表
	botList []protocol.Self
	// botId 反向索引 self -> int
	botId sync.Map
	// botRequestChan 机器人请求通道 int -> protocol.RawRequestType
	botRequestChan sync.Map
	// connect 连接
	connect OneBotV12Connect
	// impl 实现名称
	impl string
	// eventQueue 事件队列
	eventQueue chan any
	// config 配置
	config *OneBotV12Config
	// version 实现的版本
	version string
}

// OneBotV12Config OneBot V12连接配置
type OneBotV12Config struct {
	// ConnectType 连接类型
	ConnectType OneBotV12ConnectType
	// ConnectConfig 相应连接实现的配置
	ConnectConfig interface{}
}

func init() {
	_ = RegisterNewConnectType(ConnectTypeHttp, NewOneBotV12ConnectHttp)
	_ = RegisterNewConnectType(ConnectTypeWebhook, NewOneBotV12ConnectWebhook)
	_ = RegisterNewConnectType(ConnectTypeWebSocket, NewOneBotV12ConnectWebsocket)
	_ = RegisterNewConnectType(ConnectTypeWebSocketReverse, NewOneBotV12ConnectWebsocketReverse)
}

func RegisterNewConnectType(typeName OneBotV12ConnectType, registerFunc OneBotV12ConnectResisterFunc) error {
	if _, ok := OneBotV12ConnectImplementations.Load(typeName); ok {
		return ErrorConnectTypeExist
	}
	OneBotV12ConnectImplementations.Store(typeName, registerFunc)
	return nil
}

// NewOneBotV12Connect 创建连接
func NewOneBotV12Connect(config OneBotV12Config, onebotConfig protocol.OneBotConfig) (*OneBotV12, error) {
	// 创建连接
	onebot := OneBotV12{
		connectionUUID: util.GetUUID(),
		connectType:    config.ConnectType,
		botList:        make([]protocol.Self, 0),
		config:         &config,
	}
	if config.ConnectType == ConnectTypeHttp && config.ConnectConfig.(OneBotV12HttpConfig).EventEnable {
		// 开启事件队列的http连接
		if config.ConnectConfig.(OneBotV12HttpConfig).EventBufferSize <= 0 {
			onebot.eventQueue = make(chan any, 65536)
		} else {
			onebot.eventQueue = make(chan any, config.ConnectConfig.(OneBotV12HttpConfig).EventBufferSize)
		}
	}
	var err error
	if value, ok := OneBotV12ConnectImplementations.Load(config.ConnectType); !ok {
		return nil, ErrorConnectionNotFound
	} else {
		if onebot.connect, err = value.(OneBotV12ConnectResisterFunc)(onebotConfig, config.ConnectConfig); err != nil {
			return nil, err
		}
	}
	if err != nil {
		util.Logger.Warning("onebot v12 make connect error: " + err.Error())
		return nil, err
	}
	// 绑定接收函数
	onebot.connect.SetCallBackFunc(onebot.receiveMessage)
	util.Logger.Debug("onebot v12 make connect success")
	onebot.SendConnectEvent()
	return &onebot, nil
}

// AddBot 添加机器人
func (o *OneBotV12) AddBot(impl string, version string, oneBotVersion string, self protocol.Self) error {
	// 版本不匹配
	if oneBotVersion != o.GetVersion() {
		util.Logger.Warning("onebot v12 add bot error: " + ErrorConnectionNotV12.Error())
		return ErrorConnectionNotV12
	}
	// 查找self是否存在
	if _, ok := o.botId.Load(self); !ok {
		o.botId.Store(self, len(o.botList))
	} else {
		return ErrorBotExist
	}
	// 添加bot
	o.botList = append(o.botList, protocol.Self{
		PlatForm: self.PlatForm,
		UserId:   self.UserId,
	})
	o.version = version
	if impl != "" {
		o.impl = impl
	} else {
		o.impl = "github.com/FishZe/go-libonebot"
	}
	util.Logger.Debug("onebot v12 add bot success")
	o.SendStatusUpdate()
	return nil
}

// AddBotRequestChan 添加bot请求通道
func (o *OneBotV12) AddBotRequestChan(self protocol.Self, botRequestChan chan protocol.RawRequestType) error {
	if c, ok := o.botId.Load(self); ok {
		// 存进去
		o.botRequestChan.Store(c, botRequestChan)
		util.Logger.Debug("onebot v12 add bot request chan success")
		return nil
	}
	// 不存在Bot
	return ErrorBotNotExist
}

// GetVersion 获取版本号
func (*OneBotV12) GetVersion() string {
	return "12"
}

// GetUUID 获取连接的uuid
func (o *OneBotV12) GetUUID() string {
	return o.connectionUUID
}

func (o *OneBotV12) SendConnectEvent() {
	e := protocol.NewMetaEventConnect()
	e.Version.Version = o.version
	e.Version.Impl = o.impl
	e.Version.OnebotVersion = o.GetVersion()
	if eType, eid, err := protocol.EventCheck(protocol.Self{}, e); err != nil {
		util.Logger.Warning("onebot v12 check connect event error: " + err.Error())
	} else if err = o.SendEvent(e, eid, eType); err != nil {
		util.Logger.Warning("onebot v12 send connect event error: " + err.Error())
	}
}

func (o *OneBotV12) SendStatusUpdate() {
	e := protocol.NewMetaEventStatusUpdate()
	e.Status.Good = true
	for _, v := range o.botList {
		e.Status.Bots = append(e.Status.Bots, struct {
			Self   protocol.Self `json:"self"`
			Online bool          `json:"online"`
		}{v, true})
	}
	if eType, eid, err := protocol.EventCheck(protocol.Self{}, e); err != nil {
		util.Logger.Warning("onebot v12 send status update event error: " + err.Error())
	} else if err = o.SendEvent(e, eid, eType); err != nil {
		util.Logger.Warning("onebot v12 send status update event error: " + err.Error())
	}
}
