package infrastructure

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"gitlab.com/qr-through/entry/backend/config"
)

var LineBot *linebot.Client

func InitLineBot() {
	cfg := config.Config

	var err error
	LineBot, err = linebot.New(cfg.ChannelSecret, cfg.ChannelToken)
	if err != nil {
		panic(err)
	}
}
