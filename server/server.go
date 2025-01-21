package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/viogami/Gobot-vio/config"
	"github.com/viogami/Gobot-vio/tgbot"
)

func Run(port string) {
	// 默认打开http服务
	httpOn()
	// 根据配置文件创建bot
	botplatform := strings.Split(config.EnvConst.BotPlatform, ",")
	for _, v := range botplatform {
		switch v {
		case "cqhttp":
			wsOn(port)
		case "tgbot":
			//创建一个tgbot
			tgbot.CreateTgbot()
		default:
			// 默认打开cqhttp服务
			wsOn(port)
		}
	}
}

func httpOn() {
	// 设置 /post 路径的 HTTP 处理函数
	http.HandleFunc("/post", GptMsgHandle)
}

func wsOn(port string) {
	// 处理WebSocket请求的路由
	http.HandleFunc("/ws", GocqWsHandle)
	// 启动 Web 服务器监听 port 端口
	err := http.ListenAndServe(":"+port, nil)
	log.Println("HTTP server is running on port:", port)
	if err != nil {
		log.Printf("Error starting server: %v\n", err)
	}
}
