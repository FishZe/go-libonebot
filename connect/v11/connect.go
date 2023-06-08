package v11

import (
	"errors"
	"github.com/FishZe/go-libonebot/protocol"
	"github.com/FishZe/go-libonebot/util"
)

var (
	// ErrorOneBotV11ConnectedBot v11协议只能连接一个bot
	ErrorOneBotV11ConnectedBot = errors.New("OneBot V11 connected a bot")
	// ErrorConnectionNotV11 连接了错误的版本
	ErrorConnectionNotV11 = errors.New("the connection is not oneBot v11")
)

type OneBotV11 struct {
	connectType    int
	haveBot        bool
	ConnectionUUID string
}

type OneBotV11Config struct {
}

func NewOneBotV12Connect(config OneBotV11Config) (*OneBotV11, error) {
	return &OneBotV11{
		haveBot:        false,
		ConnectionUUID: util.GetUUID(),
	}, nil
}

func (o *OneBotV11) AddBotRequestChan(self protocol.Self, botRequestChan chan protocol.RawRequestType) error {
	return nil
}

func (o *OneBotV11) AddBot(impl string, version string, oneBotVersion string, self protocol.Self) error {
	if o.haveBot {
		return ErrorOneBotV11ConnectedBot
	}
	if oneBotVersion != o.GetVersion() {
		return ErrorConnectionNotV11
	}
	return nil
}

func (*OneBotV11) GetVersion() string {
	return "11"
}

func (o *OneBotV11) GetUUID() string {
	return o.ConnectionUUID
}
