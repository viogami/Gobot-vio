package gocq

import (
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/viogami/Gobot-vio/chatgpt"
	"github.com/viogami/Gobot-vio/utils"
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
func send_private_msg(conn *websocket.Conn, UserID int64, GroupID int64, message_reply string) {
	// 构建消息结构
	message_send := map[string]interface{}{
		"action": "send_private_msg",
		"params": map[string]interface{}{
			"user_id":     UserID,
			"group_id":    GroupID,
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
func send_group_msg(conn *websocket.Conn, UserID int64, GroupID int64, message_reply string) {
	cq := CQCode{
		Type: "at",
		Data: map[string]interface{}{
			"qq": fmt.Sprintf("%d", UserID),
		},
	}
	message_reply = GenerateCQCode(cq) + message_reply
	// 构建消息结构
	message_send := map[string]interface{}{
		"action": "send_group_msg",
		"params": map[string]interface{}{
			"group_id":    GroupID,
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
func send_private_img(conn *websocket.Conn, UserID int64, GroupID int64, tags []string, r18 int, num int) {
	// 调用Setu API
	setu_info := utils.Get_setu(tags, r18, num)
	if setu_info.Error != "" {
		log.Println("随机色图api调用出错:", setu_info.Error)
		return
	}
	if len(setu_info.Data) == 0 {
		log.Println("随机色图api调用出错:返回数据为空")
		send_private_msg(conn, UserID, GroupID, "该tag搜索不到图片，呜呜~")
		return
	}
	// 循环发送多张图片数据
	for i := 0; i < num; i++ {
		setu_url := setu_info.Data[i].Urls.Regular
		// 构建 message_reply 切片
		message_reply := []CQCode{
			{
				Type: "node",
				Data: map[string]interface{}{
					"name": "LV",
					"uin":  "1524175162",
					"content": []CQCode{
						{
							Type: "text",
							Data: map[string]interface{}{
								"text": fmt.Sprintf("tags:%s", tags),
							},
						},
					},
				},
			},
			{
				Type: "node",
				Data: map[string]interface{}{
					"name": "LV",
					"uin":  "1524175162",
					"content": []CQCode{
						{
							Type: "image",
							Data: map[string]interface{}{
								"file": setu_url,
							},
						},
					},
				},
			},
		}
		// 构建消息结构
		message_send := map[string]interface{}{
			"action": "send_private_forward_msg",
			"params": map[string]interface{}{
				"user_id":  UserID,
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

// 发送群聊图片
func send_group_img(conn *websocket.Conn, UserID int64, GroupID int64, tags []string, r18 int, num int) {
	// 调用Setu API
	setu_info := utils.Get_setu(tags, r18, num)
	if setu_info.Error != "" {
		log.Println("随机色图api调用出错:", setu_info.Error)
		return
	}
	if len(setu_info.Data) == 0 {
		log.Println("随机色图api调用出错:返回数据为空")
		send_group_msg(conn, UserID, GroupID, "该tag搜索不到图片，呜呜~")
		return
	}
	// 循环发送多张图片数据
	for i := 0; i < num; i++ {
		setu_url := setu_info.Data[i].Urls.Regular
		// 构建 message_reply 切片
		message_reply := []CQCode{
			{
				Type: "node",
				Data: map[string]interface{}{
					"name": "LV",
					"uin":  "1524175162",
					"content": []CQCode{
						{
							Type: "text",
							Data: map[string]interface{}{
								"text": fmt.Sprintf("tags:%s", tags),
							},
						},
					},
				},
			},
			{
				Type: "node",
				Data: map[string]interface{}{
					"name": "LV",
					"uin":  "1524175162",
					"content": []CQCode{
						{
							Type: "image",
							Data: map[string]interface{}{
								"file": setu_url,
							},
						},
					},
				},
			},
		}
		// 构建消息结构
		message_send := map[string]interface{}{
			"action": "send_group_forward_msg",
			"params": map[string]interface{}{
				"group_id": GroupID,
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

// 判断是否at我
func Atme(cq CQmsg) bool {
	CQcodes := cq.CQcodes
	for _, CQcode := range CQcodes {
		if CQcode.Type == "at" && CQcode.Data["qq"] == fmt.Sprintf("%d", receivedEvent.SelfID) {
			return true
		}
	}
	return false
}

// 请求CQ码
func GetCQCode_HuntSound(input string) string {
	sound := utils.HuntSound{
		Name:     "",
		Distance: "",
	}
	parts := strings.Split(input, " ")
	if len(parts) == 2 {
		sound.Name = parts[1]
	}
	if len(parts) == 3 {
		sound.Name = parts[1]
		sound.Distance = parts[2]
	}

	return "[CQ:record,url=" + utils.GetHuntSound(sound) + "]"
}
