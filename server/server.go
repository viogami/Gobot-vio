package server

import (
	"log/slog"
	"net/http"

	redis "github.com/redis/go-redis/v9"
	"github.com/viogami/viogo/gocq"
)

type Server struct {
	Port  string
	redis *redis.Client
	gocq  *gocq.GocqServer
}

func (s *Server) Run() {
	s.gocq = gocq.NewGocqServer(s.redis)

	// /post 处理ai请求的路由
	http.HandleFunc("/post", gptMsgHandle)
	// 处理WebSocket请求的路由
	http.HandleFunc("/ws", handleWebSocket)

	// 启动 Web 服务器监听 port 端口
	err := http.ListenAndServe(":"+s.Port, nil)
	slog.Info("Server started", "port", s.Port)
	if err != nil {
		slog.Error("Error starting server:", "err", err)
	}
}

func (s *Server) WithRedis(redisURL string) {
	if redisURL == "" {
		slog.Error("REDIS_URL environment variable not set")
	}
	// 解析 Redis URL
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		slog.Error("Failed to parse Redis URL:", "err", err)
	}
	s.redis = redis.NewClient(opt)
}

func NewServer(port string) *Server {
	return &Server{
		Port: port,
		redis: new(redis.Client),
		gocq:  new(gocq.GocqServer),
	}
}
