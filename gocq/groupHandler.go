package gocq

import (
	"log"

	"github.com/gorilla/websocket"
)

func set_group_ban(conn *websocket.Conn, MsgEvent *MessageEvent, time int) {
	// 构建消息结构
	sendMessage := map[string]interface{}{
		"action": "set_group_ban",
		"params": map[string]interface{}{
			"group_id": MsgEvent.GroupID,
			"user_id":  MsgEvent.UserID,
			"duration": time, // 单位秒，0 表示解除禁言
		},
		"echo": "echo_test",
	}
	// 发送 JSON 消息
	err := conn.WriteJSON(sendMessage)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}
