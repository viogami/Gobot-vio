package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/viogami/Gobot-vio/gocq"

	"github.com/gorilla/websocket"
)

// handleWebSocket 用于处理WebSocket请求
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		// 从WebSocket连接读取消息
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		// 打印接收到的消息
		err = gocq.Log_post_type(p)
		if err != nil {
			log.Println(err)
		} else {
			// 发送消息
			gocq.Handle_event(p, conn)
		}
	}
}
