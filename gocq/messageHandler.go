package gocq

import (
	"Gobot-vio/chatgpt"
	"Gobot-vio/utils"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// 消息处理函数
func msgHandler(MsgEvent *MessageEvent) string {
	msgText := ParseCQmsg(MsgEvent.Message).Text

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
func send_private_msg(conn *websocket.Conn, MsgEvent *MessageEvent, message_reply string) {
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
func send_group_msg(conn *websocket.Conn, MsgEvent *MessageEvent, message_reply string) {
	cq := CQCode{
		Type: "at",
		Data: map[string]interface{}{
			"qq":   fmt.Sprintf("%d", MsgEvent.UserID),
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

// 发送图片
func send_private_img(conn *websocket.Conn, MsgEvent *MessageEvent, tags []string, r18 int, num int) {
	// 调用Setu API
	setu_info := utils.Get_setu(tags, r18, num)
	if setu_info.Error != "" {
		log.Println("随机色图api调用出错:", setu_info.Error)
		return
	}
	if len(setu_info.Data) == 0 {
		log.Println("随机色图api调用出错:返回数据为空")
		send_private_msg(conn, MsgEvent, "该tag搜索不到图片，呜呜~")
		return
	}
	// 循环发送多张图片数据
	for i := 0; i < num; i++ {
		setu_url := setu_info.Data[i].Urls.Regular
		// 图片cq码
		cq := CQCode{
			Type: "image",
			Data: map[string]interface{}{
				"file": setu_url,
			},
		}
		message_reply := []CQCode{
			{
				Type: "node",
				Data: map[string]interface{}{
					"name":    "LV",
					"uin":     "1524175162",
					"content": fmt.Sprintf("涩图 tags:%s", tags),
				},
			}, {
				Type: "node",
				Data: map[string]interface{}{
					"name":    "LV",
					"uin":     "1524175162",
					"content": GenerateCQCode(cq),
				},
			},
		}
		log.Println(message_reply)
		// // 构建消息结构
		// message_send := map[string]interface{}{
		// 	"action": "send_private_forward_msg",
		// 	"params": map[string]interface{}{
		// 		"user_id":  MsgEvent.UserID,
		// 		"messages": message_reply,
		// 	},
		// 	"echo": "echo_test",
		// }

		///////////////////////////////////////////////
		message_reply2 := GenerateCQCode(cq)
		// 构建消息结构
		message_send2 := map[string]interface{}{
			"action": "send_msg",
			"params": map[string]interface{}{
				"message_type": MsgEvent.MessageType,
				"user_id":      MsgEvent.UserID,
				"group_id":     MsgEvent.GroupID,
				"message":      message_reply2,
				"auto_escape":  false,
			},
			"echo": "echo_test",
		}
		// 发送 JSON 消息
		err := conn.WriteJSON(message_send2)
		if err != nil {
			log.Println("Error sending message:", err)
		}
	}
}

// 发送群聊图片
func send_group_img(conn *websocket.Conn, MsgEvent *MessageEvent, tags []string, r18 int, num int) {
	// 调用Setu API
	setu_info := utils.Get_setu(tags, r18, num)
	if setu_info.Error != "" {
		log.Println("随机色图api调用出错:", setu_info.Error)
		return
	}
	if len(setu_info.Data) == 0 {
		log.Println("随机色图api调用出错:返回数据为空")
		send_group_msg(conn, MsgEvent, "该tag搜索不到图片，呜呜~")
		return
	}
	// 循环发送多张图片数据
	for i := 0; i < num; i++ {
		setu_url := setu_info.Data[i].Urls.Regular
		// 图片cq码
		cq := CQCode{
			Type: "image",
			Data: map[string]interface{}{
				"file": setu_url,
			},
		}
		message_reply := []CQCode{
			{
				Type: "node",
				Data: map[string]interface{}{
					"name":    "LV",
					"uin":     "1524175162",
					"content": fmt.Sprintf("涩图 tags:%s", tags),
				},
			}, {
				Type: "node",
				Data: map[string]interface{}{
					"name":    "LV",
					"uin":     "1524175162",
					"content": GenerateCQCode(cq),
				},
			},
		}
		log.Println(message_reply)
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
}