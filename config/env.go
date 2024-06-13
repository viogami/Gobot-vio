package config

import "os"

type ENV struct {
	PORT             string
	BOT_TOKEN        string
	TG_WEBHOOK_URL   string
	ChatGPTAPIKey    string
	ChatGPTURL_proxy string
	BotPlatform      string
}

var EnvConst ENV

func GetEnv() ENV {
	EnvConst.PORT = os.Getenv("PORT")
	EnvConst.BOT_TOKEN = os.Getenv("BOT_TOKEN")
	EnvConst.TG_WEBHOOK_URL = os.Getenv("TG_WEBHOOK_URL")
	EnvConst.ChatGPTAPIKey = os.Getenv("ChatGPTAPIKey")
	EnvConst.ChatGPTURL_proxy = os.Getenv("ChatGPTURL_proxy")
	EnvConst.BotPlatform = os.Getenv("BotPlatform")

	return EnvConst
}
