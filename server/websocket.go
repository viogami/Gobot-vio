package server

import (
	"Gobot-vio/gocq"
	"fmt"
	"net/http"

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
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// 打印接收到的消息
		fmt.Print("messageType:", messageType)
		fmt.Printf("Received message: %s\n", p)

		// 发送消息
		gocq.Send_private_msg(conn, 2654613995, string(p))
	}
}
