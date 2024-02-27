package server

import (
	"log"
	"net/http"

	"github.com/viogami/Gobot-vio/tgbot"
)

func Run(port string) {
	// 设置 /post 路径的 HTTP 处理函数
	http.HandleFunc("/post", handlePost)
	// 处理WebSocket请求的路由
	http.HandleFunc("/ws", handleWebSocket)
	// 启动 Web 服务器监听 port 端口
	go func() {
		err := http.ListenAndServe(":"+port, nil)
		log.Println("HTTP server is running on port:", port)
		if err != nil {
			log.Printf("Error starting server: %v\n", err)
		}
	}()

	//创建一个tgbot
	tgbot.CreateTgbot()
}
