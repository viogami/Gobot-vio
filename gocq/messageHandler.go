package gocq

import (
	"Gobot-vio/chatgpt"
	"Gobot-vio/utils"
	"log"

	"github.com/gorilla/websocket"
)

// 消息处理函数
func msgHandler(MsgEvent *MessageEvent) string {
	msgText := ParseCQmsg(receivedMsgEvent.Message).Text

	reply_res := Msg_Filter(msgText)

	if reply_res == "" {
		log.Println("调用ChatGPT API")
		gptResponse, err := chatgpt.InvokeChatGPTAPI(msgText)
		if err != nil {
			log.Printf("Error calling ChatGPT API: %v", err)
		}
		return gptResponse
	}
	return reply_res
}

// 发送私聊消息
func send_private_msg(conn *websocket.Conn, MsgEvent *MessageEvent) {
	message_reply := msgHandler(MsgEvent)
	// 构建消息结构
	message_send := map[string]interface{}{
		"action": "send_private_msg",
		"params": map[string]interface{}{
			"user_id":     MsgEvent.UserID,
			"group_id":    MsgEvent.GroupID,
			"message":     message_reply,
			"auto_escape": false, // 消息内容是否作为纯文本发送 ( 即不解析 CQ 码 )，只在 message 字段是字符串时有效
		},
		"echo": "echo_test", // 用于识别回调消息
	}
	// 发送 JSON 消息
	err := conn.WriteJSON(message_send)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

// 发送群聊消息
func send_group_msg(conn *websocket.Conn, MsgEvent *MessageEvent) {
	message_reply := msgHandler(MsgEvent)
	cq := CQCode{
		Type: "at",
		Params: map[string]interface{}{
			"qq":   MsgEvent.UserID,
			"name": "不在群的QQ",
		},
	}
	message_reply = GenerateCQCode(cq) + message_reply
	// 构建消息结构
	message_send := map[string]interface{}{
		"action": "send_group_msg",
		"params": map[string]interface{}{
			"group_id":    MsgEvent.GroupID,
			"message":     message_reply,
			"auto_escape": false,
		},
		"echo": "echo_test",
	}
	// 发送 JSON 消息
	err := conn.WriteJSON(message_send)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

// TODO: 发送群聊合并转发消息
func send_group_forward_msg(conn *websocket.Conn, MsgEvent *MessageEvent) {
	message_reply := msgHandler(MsgEvent)
	cq := CQCode{
		Type: "at",
		Params: map[string]interface{}{
			"qq":   MsgEvent.UserID,
			"name": "不在群的QQ",
		},
	}
	message_reply = GenerateCQCode(cq) + message_reply
	// 构建消息结构
	message_send := map[string]interface{}{
		"action": "send_group_forward_msg",
		"params": map[string]interface{}{
			"group_id": MsgEvent.GroupID,
			"messages": message_reply,
		},
		"echo": "echo_test",
	}
	// 发送 JSON 消息
	err := conn.WriteJSON(message_send)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

// 发送图片
func send_image(conn *websocket.Conn, MsgEvent *MessageEvent, tags []string, r18 int, num int) {
	// 调用Setu API
	setu_info := utils.Get_setu(tags, r18, num)
	if setu_info.Error != "" {
		log.Println("随机色图api调用出错:", setu_info.Error)
		return
	}
	// 循环发送多张图片数据
	for i := 0; i < num; i++ {
		setu_url := setu_info.Data[i].Urls.Regular
		cq := CQCode{
			Type: "image",
			Params: map[string]interface{}{
				"url": setu_url,
			},
		}
		message_reply := GenerateCQCode(cq)
		// 构建消息结构
		message_send := map[string]interface{}{
			"action": "send_msg",
			"params": map[string]interface{}{
				"message_type": MsgEvent.MessageType,
				"user_id":      MsgEvent.UserID,
				"group_id":     MsgEvent.GroupID,
				"message":      message_reply,
				"auto_escape":  false,
			},
			"echo": "echo_test",
		}
		// 发送 JSON 消息
		err := conn.WriteJSON(message_send)
		if err != nil {
			log.Println("Error sending message:", err)
		}
	}

}
