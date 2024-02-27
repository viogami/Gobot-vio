package gocq

import (
	"log"

	"github.com/gorilla/websocket"
)

// 设置群组禁言
func set_group_ban(conn *websocket.Conn, UserID int64, GroupID int64, time int) {
	// 构建消息结构
	sendMessage := map[string]interface{}{
		"action": "set_group_ban",
		"params": map[string]interface{}{
			"group_id": GroupID,
			"user_id":  UserID,
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

// 撤回消息
func delete_msg(conn *websocket.Conn, MsgEvent *MessageEvent, messageID int) {
	// 构建消息结构
	sendMessage := map[string]interface{}{
		"action": "delete_msg",
		"params": map[string]interface{}{
			"message_id": messageID,
		},
		"echo": "echo_test",
	}
	// 发送 JSON 消息
	err := conn.WriteJSON(sendMessage)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}
