package v11

import (
	"github.com/FishZe/go-libonebot/util"
	"log"
)

func (o *OneBotV11) ConnectSendEvent(e any, eId string) error {
	s, err := util.Json.Marshal(e)
	if err != nil {
		return err
	}
	log.Println(string(s))
	return nil
}

func (o *OneBotV11) ConnectSendResponse(e any) error {
	return nil
}
