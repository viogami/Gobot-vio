package server

import (
	"log"
	"net/http"

	redis "github.com/go-redis/redis/v8"

	"github.com/viogami/Gobot-vio/tgbot"
)

type Server struct {
	Port     string
	redis    *redis.Client
}

func (s *Server) httpOn() {
	// 设置 /post 路径的 HTTP 处理函数
	http.HandleFunc("/post", GptMsgHandle)
}

func (s *Server) wsOn(port string) {
	// 处理WebSocket请求的路由
	http.HandleFunc("/ws", GocqWsHandle)
	// 启动 Web 服务器监听 port 端口
	err := http.ListenAndServe(":"+port, nil)
	log.Println("HTTP server is running on port:", port)
	if err != nil {
		log.Printf("Error starting server: %v\n", err)
	}
}

func (s *Server) tgbotOn() {
	// 创建一个tgbot
	tgbot.CreateTgbot()
}

func Run(port string) {
	s := new(Server)
	s.Port = port
	// s.redis = redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379", // Redis 服务器地址和端口
	// 	Password: "",               // Redis 密码，如果没有设置则为空
	// 	DB:       0,               // Redis 数据库索引
	// })

	// 启动 HTTP 服务器
	s.httpOn()
	// 启动 WebSocket 服务器
	s.wsOn(port)
	// 启动 Telegram Bot
	// s.tgbotOn()
}
