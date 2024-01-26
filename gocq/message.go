package gocq

import (
	"Gobot-vio/chatgpt"
	"log"

	"github.com/gorilla/websocket"
)

func Send_msg(conn *websocket.Conn, msgtype string, targetID int64, message string) {
	// chatgpt回复
	message = reply(message)
	// 构建消息结构
	sendMessage := map[string]interface{}{
		"action": "send_msg",
		"params": map[string]interface{}{
			"user_id": targetID,
			"message": message,
		},
		"echo": "echo_test", // 用于识别回调消息
	}
	// 判断消息类型
	if msgtype == "private" {
		sendMessage["action"] = "send_private_msg"
	} else if msgtype == "group" {
		sendMessage["action"] = "send_group_msg"
	}
	// 发送 JSON 消息
	err := conn.WriteJSON(sendMessage)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func reply(text string) string {
	// 调用ChatGPT API
	gptResponse, err := chatgpt.InvokeChatGPTAPI(text)
	if err != nil {
		log.Printf("Error calling ChatGPT API: %v", err)
		gptResponse = "gpt调用失败了😥 错误信息：\n" + err.Error()
	}
	return gptResponse
}
