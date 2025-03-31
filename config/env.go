package config

import "os"

var EnvConst ENV

type ENV struct {
	PORT             string

	ENV_AI
	ENV_TG
}
type ENV_AI struct {
	ChatGPTAPIKey    string
	ChatGPTURL_proxy string
}
type ENV_TG struct {
	BOT_TOKEN      string
	TG_WEBHOOK_URL string
}

func GetEnv() ENV {
	EnvConst.PORT = os.Getenv("PORT")

	EnvConst.ChatGPTAPIKey = os.Getenv("ChatGPTAPIKey")
	EnvConst.ChatGPTURL_proxy = os.Getenv("ChatGPTURL_proxy")

	EnvConst.BOT_TOKEN = os.Getenv("BOT_TOKEN")
	EnvConst.TG_WEBHOOK_URL = os.Getenv("TG_WEBHOOK_URL")

	return EnvConst
}
