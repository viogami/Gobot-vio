package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Tgbot   tgbotConfig   `yaml:"tgbot"`
	Chatgpt chatgptConfig `yaml:"chatgpt"`
	Wx      wxConfig      `yaml:"wx"`
}

type ServerConfig struct {
	Env     string `yaml:"env"`
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type tgbotConfig struct {
	TG_WEBHOOK_URL string `yaml:"TG_WEBHOOK_URL"`
	BOT_TOKEN      string `yaml:"BOT_TOKEN"`
}

type chatgptConfig struct {
	ChatGPTAPIKey   string `yaml:"chatGPTAPIKey"`
	ChatGPTURL_chat string `yaml:"chatGPTURL_chat"`
	ChatGPTURL_img  string `yaml:"chatGPTURL_img"`
	ChatGPTURL_mood string `yaml:"chatGPTURL_mood"`
}

type wxConfig struct {
	Port    int    `yaml:"port"`
	WxTOKEN string `yaml:"wxTOKEN"`
}

// 读取配置文件 app.yaml
func ConfigParse(appConfig string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(appConfig)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
