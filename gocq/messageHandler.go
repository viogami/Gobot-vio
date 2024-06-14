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
			gptResponse = "Error calling ChatGPT API" + fmt.Sprintf("%v", err)
		}
		return gptResponse
	}
	return reply_res
}

// 发送私聊消息
// @params
// message_type	string	-	消息类型, 支持 private、group , 分别对应私聊、群组, 如不传入, 则根据传入的 *_id 参数判断
// user_id	int64	-	对方 QQ 号 ( 消息类型为 private 时需要 )
// group_id	int64	-	群号 ( 消息类型为 group 时需要 )
// message	message	-	要发送的内容
// auto_escape	boolean	false	消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 message 字段是字符串时有效
func send_msg(conn *websocket.Conn, message_type string, user_id int64, group_id int64, message string, auto_escape bool) {
	if message_type == "group" {
		cq := CQCode{
			Type: "at",
			Data: map[string]interface{}{
				"qq": fmt.Sprintf("%d", receivedMsgEvent.UserID),
			},
		}
		message = GenerateCQCode(cq) + message
	}
	// 构建消息结构
	message_send := map[string]interface{}{
		"action": "send_msg",
		"params": map[string]interface{}{
			"message_type": message_type,
			"user_id":      user_id,
			"group_id":     group_id,
			"message":      message,
			"auto_escape":  auto_escape,
		},
		"echo": "echo_test", // 用于识别回调消息
	}
	// 发送 JSON 消息
	err := conn.WriteJSON(message_send)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

// 发送私聊合并消息
func send_private_forward_msg(conn *websocket.Conn, UserID int64, message_reply []CQCode) {
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
// 发送群聊合并消息
func send_group_forward_msg(conn *websocket.Conn, GroupID int64, message_reply []CQCode) {
		// 构建消息结构
		message_send := map[string]interface{}{
			"action": "send_group_forward_msg",
			"params": map[string]interface{}{
				"group_id":  GroupID,
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

	return "[CQ:record,file=" + utils.GetHuntSound(sound) + "]"
}
