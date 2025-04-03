package config

type ENV struct {
	PORT string
}
type ENV_AI struct {
	ChatGPTAPIKey    string
	ChatGPTURL_proxy string
}
type ENV_TG struct {
	BOT_TOKEN      string
	TG_WEBHOOK_URL string
}
type ENV_REDIS struct {
	REDIS_URL string
}
