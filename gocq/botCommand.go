package gocq

import (
	"fmt"
	"log"
	"math/rand"

	group "github.com/viogami/Gobot-vio/gocq/groupManage"
	"github.com/viogami/Gobot-vio/utils"
)

type cmd_params struct {
	receivedMsgEvent *MessageEvent
	tags             []string
	num              int
}

var privateCommandList = map[string]func(params cmd_params) map[string]interface{}{
	"":       privateCmd_null, // 空指令，不做任何处理
	"/help":  privateCmd_help,
	"/涩图":    privateCmd_setu,
	"/涩图r18": privateCmd_r18,
	"/枪声":    privateCmd_HuntSound,
	"/枪声目录":  privateCmd_HuntSoundList,
}
var groupCommandList = map[string]func(cmd_params) map[string]interface{}{
	"":       groupCmd_null, // 空指令，不做任何处理
	"/help":  groupCmd_help,
	"/chat":  groupCmd_chat,
	"/涩图":    groupCmd_setu,
	"/涩图r18": groupCmd_r18,
	"/枪声":    groupCmd_HuntSound,
	"/枪声目录":  groupCmd_HuntSoundList,
	"/禁言抽奖":  groupCmd_BanLottery,
}

const (
	privateCmd = "私聊指令:\n" + "/help:查看帮助\n" + "/涩图:随机涩图\n" + "/涩图r18:随机r18涩图\n" + "/枪声:随机枪声\n" + "/枪声目录:枪声目录\n"
	groupCmd   = "群聊指令:\n" + "/help:查看帮助\n" + "/chat:聊天\n" + "/涩图:随机涩图\n" + "/涩图r18:随机r18涩图\n" + "/枪声:随机枪声\n" + "/枪声目录:枪声目录\n" + "/禁言抽奖:禁言抽奖0~180秒\n"
)

// ---------私聊指令处理函数---------
// 空指令
func privateCmd_null(params cmd_params) map[string]interface{} {
	receivedMsgEvent := params.receivedMsgEvent

	log.Printf("将对私聊回复,msgID:%d,UserID:%d", receivedMsgEvent.MessageID, receivedMsgEvent.UserID)
	// 消息处理
	message_reply := msgGptHandler(receivedMsgEvent)
	return msg_send(receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, message_reply, false)
}

// help 指令
func privateCmd_help(params cmd_params) map[string]interface{} {
	receivedMsgEvent := params.receivedMsgEvent

	log.Printf("将对私聊回复,msgID:%d,UserID:%d", receivedMsgEvent.MessageID, receivedMsgEvent.UserID)
	return msg_send(receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, privateCmd, false)
}

// 涩图 指令
func privateCmd_setu(params cmd_params) map[string]interface{} {
	receivedMsgEvent := params.receivedMsgEvent
	tags := params.tags
	num := params.num
	UserID := receivedMsgEvent.UserID
	GroupID := receivedMsgEvent.GroupID

	log.Printf("将对私聊发送涩图 tags:%s,nums:%d", tags, num)
	// 得到色图消息
	message_reply := get_setu_MsgReply(tags, 0, num)
	if message_reply == nil {
		msg_send(receivedMsgEvent.MessageType, UserID, GroupID, "涩图获取失败,tag搜索不到图片...", false)
	}
	// 发送消息
	return msg_send_private_forward(UserID, message_reply)
}

// r18 指令
func privateCmd_r18(params cmd_params) map[string]interface{} {
	receivedMsgEvent := params.receivedMsgEvent
	tags := params.tags
	num := params.num
	UserID := receivedMsgEvent.UserID
	GroupID := receivedMsgEvent.GroupID

	log.Printf("将对私聊发送r18涩图 tags:%s,nums:%d", tags, num)
	// 得到色图消息
	message_reply := get_setu_MsgReply(tags, 1, num)
	if message_reply == nil {
		msg_send(receivedMsgEvent.MessageType, UserID, GroupID, "涩图获取失败,tag搜索不到图片...", false)
	}
	// 发送消息
	return msg_send_private_forward(UserID, message_reply)
}

// 枪声 指令
func privateCmd_HuntSound(params cmd_params) map[string]interface{} {
	receivedMsgEvent := params.receivedMsgEvent

	// 解析cq码，获取无cq格式的消息内容
	cqmsg := ParseCQmsg(receivedMsgEvent.Message)
	log.Println("将对私聊发送枪声:" + GetCQCode_HuntSound(cqmsg.Text))
	return msg_send(receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, GetCQCode_HuntSound(cqmsg.Text), false)
}

// 枪声目录 指令
func privateCmd_HuntSoundList(params cmd_params) map[string]interface{} {
	receivedMsgEvent := params.receivedMsgEvent

	log.Println("将对私聊发送枪声目录")
	return msg_send(receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, utils.GetGunIndex(), false)
}

// ---------群聊指令处理函数---------
// null 指令
func groupCmd_null(params cmd_params) map[string]interface{} {
	log.Printf("非指令消息,sender:%d,groupid:%d", params.receivedMsgEvent.Sender.UserID, params.receivedMsgEvent.GroupID)
	return nil
}

// help 指令
func groupCmd_help(params cmd_params) map[string]interface{} {
	receivedMsgEvent := params.receivedMsgEvent
	log.Printf("将对群聊回复,msgID:%d,UserID:%d,GroupID:%d", receivedMsgEvent.MessageID, receivedMsgEvent.UserID, receivedMsgEvent.GroupID)
	return msg_send(receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, groupCmd, false)
}

// chat 指令
func groupCmd_chat(params cmd_params) map[string]interface{} {
	receivedMsgEvent := params.receivedMsgEvent

	log.Printf("将对群聊回复,msgID:%d,UserID:%d,GroupID:%d", receivedMsgEvent.MessageID, receivedMsgEvent.UserID, receivedMsgEvent.GroupID)
	// 消息处理
	message_reply := msgGptHandler(receivedMsgEvent)
	return msg_send(receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, message_reply, false)
}

// 涩图 指令
func groupCmd_setu(params cmd_params) map[string]interface{} {
	receivedMsgEvent := params.receivedMsgEvent
	tags := params.tags
	num := params.num
	UserID := receivedMsgEvent.UserID
	GroupID := receivedMsgEvent.GroupID

	log.Printf("将对群聊发送涩图 tags:%s,nums:%d", tags, num)
	// 得到色图消息
	message_reply := get_setu_MsgReply(tags, 0, num)
	if message_reply == nil {
		msg_send(receivedMsgEvent.MessageType, UserID, GroupID, "涩图获取失败,tag搜索不到图片...", false)
	}
	// 返回将要发送的消息
	return msg_send_group_forward(GroupID, message_reply)
}

// r18 指令
func groupCmd_r18(params cmd_params) map[string]interface{} {
	receivedMsgEvent := params.receivedMsgEvent
	tags := params.tags
	num := params.num
	UserID := receivedMsgEvent.UserID
	GroupID := receivedMsgEvent.GroupID

	log.Printf("将对群聊发送r18涩图 tags:%s,nums:%d", tags, num)
	// 得到色图消息
	message_reply := get_setu_MsgReply(tags, 1, num)
	if message_reply == nil {
		msg_send(receivedMsgEvent.MessageType, UserID, GroupID, "涩图获取失败,tag搜索不到图片...", false)
	}
	// 返回将要发送的消息
	return msg_send_group_forward(GroupID, message_reply)
}

// 枪声 指令
func groupCmd_HuntSound(params cmd_params) map[string]interface{} {
	receivedMsgEvent := params.receivedMsgEvent

	// 解析cq码，获取无cq格式的消息内容'
	cqmsg := ParseCQmsg(receivedMsgEvent.Message)
	log.Println("将对群聊发送枪声" + GetCQCode_HuntSound(cqmsg.Text))
	return msg_send(receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, GetCQCode_HuntSound(cqmsg.Text), false)
}

// 枪声目录 指令
func groupCmd_HuntSoundList(params cmd_params) map[string]interface{} {
	receivedMsgEvent := params.receivedMsgEvent

	log.Println("将对群聊发送枪声目录")
	return msg_send(receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, utils.GetGunIndex(), false)
}

// 禁言抽奖 指令
func groupCmd_BanLottery(params cmd_params) map[string]interface{} {
	receivedMsgEvent := params.receivedMsgEvent
	time := rand.Intn(180) + 1
	log.Printf("将对群聊:%d,禁言qq用户:%d,时间:%d秒", receivedMsgEvent.GroupID, receivedMsgEvent.UserID, time)
	replyMsgs = append(replyMsgs, msg_send("group", receivedMsgEvent.UserID, receivedMsgEvent.GroupID, fmt.Sprintf("恭喜中奖，禁言时间:%d秒~", time), false))
	return group.Set_group_ban(receivedMsgEvent.UserID, receivedMsgEvent.GroupID, time)
}

// ------------------------------- 可复用代码 ----------------------------------
// 寻找色图
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
