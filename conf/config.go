package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

var AppConfig Config

type Config struct {
	EnabledService
	AdminPanel
}

type EnabledService struct {
	AdminPanelEnabled bool `yaml:"admin_panel_enabled"`
	RedisEnabled      bool `yaml:"redis_enabled"`
	MongoDBEnabled    bool `yaml:"mongodb_enabled"`
}

type AdminPanel struct {
	AdminPanelHost string `yaml:"host"`
	AdminPanelPort string `yaml:"port"`
	AdminPanelUser string `yaml:"username"`
	AdminPanelPass string `yaml:"password"`
}

// 初始化配置
func Init() {
	// 读取配置文件
	data, err := os.ReadFile("conf/config.yaml")
	if err != nil {
		slog.Error("读取配置文件失败", "error", err)
	}

	// 解析 YAML
	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		slog.Error("解析配置文件失败", "error", err)
	}
}
