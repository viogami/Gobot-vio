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
			"message_type": msgtype,  // "private" / "group
			"user_id":      targetID, // 仅在发送私聊消息时使用
			"group_id":     targetID, // 仅在发送群消息时使用
			"message":      message,
			"auto_escape":  false, // 消息内容是否作为纯文本发送 ( 即不解析 CQ 码 )，只在 message 字段是字符串时有效
		},
		"echo": "echo_test", // 用于识别回调消息
	}
	// 判断消息类型
	if msgtype != "private" && msgtype != "group" {
		log.Println("Error: msgtype is not private or group")
		return
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
