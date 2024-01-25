package server

import (
	"log"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Run(port string) {
	// 设置 /post 路径的 HTTP 处理函数
	s.mux.HandleFunc("/post", handlePost)
	// 处理WebSocket请求的路由
	s.mux.HandleFunc("/ws", handleWebSocket)
	// 启动 Web 服务器监听 port 端口
	go func() {
		err := http.ListenAndServe(":"+port, s.mux)
		log.Println("HTTP server is running on port:", port)
		if err != nil {
			log.Printf("Error starting server: %v\n", err)
		}
	}()
}

// Close 方法用于关闭服务器
func (s *Server) Close() {

}
