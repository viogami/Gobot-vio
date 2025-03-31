package gocq

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/viogami/Gobot-vio/utils"
)

// 消息处理函数
func msgGptHandler(MsgEvent *MessageEvent) string {
	msgText := ParseCQmsg(MsgEvent.Message).Text

	reply_res := utils.Msg_Filter(msgText)

	if reply_res == "" {
		log.Println("调用ChatGPT API")
		gptResponse, err := "nil", errors.New("nil")
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
func msg_send(message_type string, user_id int64, group_id int64, message string, auto_escape bool) map[string]interface{} {
	if message_type == "group" {
		cq := CQCode{
			Type: "at",
			Data: map[string]interface{}{
				"qq": fmt.Sprintf("%d", user_id),
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
	return message_send
}

// 发送私聊合并消息
func msg_send_private_forward(UserID int64, message_reply []CQCode) map[string]interface{} {
	// 构建消息结构
	message_send := map[string]interface{}{
		"action": "send_private_forward_msg",
		"params": map[string]interface{}{
			"user_id":  UserID,
			"messages": message_reply,
		},
		"echo": "echo_test",
	}
	return message_send
}

// 发送群聊合并消息
func msg_send_group_forward(GroupID int64, message_reply []CQCode) map[string]interface{} {
	// 构建消息结构
	message_send := map[string]interface{}{
		"action": "send_group_forward_msg",
		"params": map[string]interface{}{
			"group_id": GroupID,
			"messages": message_reply,
		},
		"echo": "echo_test",
	}
	return message_send
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

// 获得猎杀枪声的CQ码
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

// 构建涩图消息回复
func get_setu_MsgReply(tags []string, r18 int, num int) []CQCode {
	// 调用Setu API
	setu_info := utils.Get_setu(tags, r18, num)
	if setu_info.Error != "" {
		log.Println("随机色图api调用出错:", setu_info.Error)
		return nil
	}
	if len(setu_info.Data) == 0 {
		log.Println("随机色图api调用出错:tag搜索不到，返回数据为空")
		return nil
	}
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
	}
	// 循环存放多张图片数据
	for i := 0; i < num; i++ {
		setu_url := setu_info.Data[i].Urls.Regular
		// 构建 message_reply 切片
		setu_cqNode := CQCode{
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
		}
		message_reply = append(message_reply, setu_cqNode)
	}
	return message_reply
}
