package v12

import (
	"errors"
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
	"sync"
)

const (
	// ConnectTypeHttp http连接
	ConnectTypeHttp = 1 << 0
	// ConnectTypeWebhook webhook 连接
	ConnectTypeWebhook = 1 << 1
	// ConnectTypeWebSocket websocket连接
	ConnectTypeWebSocket = 1 << 2
	// ConnectTypeWebSocketReverse 反向websocket连接
	ConnectTypeWebSocketReverse = 1 << 3
)

const (
	sendTypeResponse = 1 << 0
	SendTypeEvent    = 1 << 1
)

var (
	// ErrorConnectionNotV12 连接了错误的版本
	ErrorConnectionNotV12 = errors.New("the connection is not oneBot v12")
	// ErrorBotExist bot已经存在
	ErrorBotExist = errors.New("bot already exists")
	// ErrorBotNotExist bot不存在
	ErrorBotNotExist = errors.New("bot does not exist")
	// ErrorTooMuchConnection 连接过多
	ErrorTooMuchConnection = errors.New("too much connection")
	// ErrorNoConnection 没有连接
	ErrorNoConnection = errors.New("no connection")
)

// OneBotV12Connect 连接接口
type OneBotV12Connect interface {
	Send([]byte, int, string) error
	BindReceiveFunc(func([]byte) (string, error))
}

// OneBotV12 OneBot V12连接
type OneBotV12 struct {
	// 连接类型
	connectType int
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
}

// OneBotV12Config OneBot V12连接配置
type OneBotV12Config struct {
	// ConnectType 连接类型
	ConnectType int
	// HttpConfig http连接配置
	HttpConfig OneBotV12HttpConfig
	// WebhookConfig webhook连接配置
	WebhookConfig OneBotV12WebHookConfig
	// WebSocketConfig websocket连接配置
	WebsocketConfig OneBotV12WebsocketConfig
	// WebSocketReverseConfig 反向websocket连接配置
	WebsocketReverseConfig OneBotV12WebsocketReverseConfig
}

// NewOneBotV12Connect 创建连接
func NewOneBotV12Connect(config OneBotV12Config) (*OneBotV12, error) {
	// 判断连接数量
	var count = 0
	var connectType = config.ConnectType
	for connectType != 0 {
		connectType = connectType & (connectType - 1)
		count++
	}
	// 数量不为1
	if count == 0 {
		return nil, ErrorNoConnection
	} else if count > 1 {
		return nil, ErrorTooMuchConnection
	}
	// 创建连接
	onebot := OneBotV12{
		connectionUUID: util.GetUUID(),
		botList:        make([]protocol.Self, 0),
	}
	var err error
	switch config.ConnectType {
	case ConnectTypeHttp:
		// HTTP
		onebot.connect, err = NewOneBotV12ConnectHttp(&config.HttpConfig)
	case ConnectTypeWebhook:
		// Webhook
		config.WebhookConfig.impl = onebot.impl
		onebot.connect, err = NewOneBotV12ConnectWebhook(&config.WebhookConfig)
	case ConnectTypeWebSocket:
		// Websocket 作为websocket server
		config.WebsocketReverseConfig.impl = onebot.impl
		onebot.connect, err = NewOneBotV12ConnectWebsocket(&config.WebsocketConfig)
	case ConnectTypeWebSocketReverse:
		// WebsocketReverse 作为websocket client
		onebot.connect, err = NewOneBotV12ConnectWebsocketReverse(&config.WebsocketReverseConfig)
	default:
		err = ErrorNoConnection
	}
	if err != nil {
		return nil, err
	}
	// 绑定接收函数
	onebot.connect.BindReceiveFunc(onebot.receiveMessage)
	return &onebot, nil
}

// AddBot 添加机器人
func (o *OneBotV12) AddBot(impl string, version string, oneBotVersion string, self protocol.Self) error {
	// 版本不匹配
	if oneBotVersion != o.GetVersion() {
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
	o.impl = impl
	return nil
}

// AddBotRequestChan 添加bot请求通道
func (o *OneBotV12) AddBotRequestChan(self protocol.Self, botRequestChan chan protocol.RawRequestType) error {
	if c, ok := o.botId.Load(self); ok {
		// 存进去
		o.botRequestChan.Store(c, botRequestChan)
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
