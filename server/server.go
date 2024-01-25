package server

import (
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

func (s *Server) Run(port string) error {
	// 设置 / 路径
	s.mux.HandleFunc("/", handleHome)
	// 设置 /post 路径的 HTTP 处理函数
	s.mux.HandleFunc("/post", handlePost)
	// 启动 Web 服务器监听 port 端口
	err := http.ListenAndServe(":"+port, s.mux)
	return err
}

// Close 方法用于关闭服务器
func (s *Server) Close() {

}
