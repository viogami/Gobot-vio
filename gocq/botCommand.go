package gocq

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gorilla/websocket"
	"github.com/viogami/Gobot-vio/utils"
)

type cmd_params struct {
	conn             *websocket.Conn
	receivedMsgEvent MessageEvent
	tags             []string
	num              int
}

var privateCommandList = map[string]func(params cmd_params){
	"":       privateCmd_null, // 空指令，不做任何处理
	"/help":  privateCmd_help,
	"/涩图":    privateCmd_setu,
	"/涩图r18": privateCmd_r18,
	"/枪声":    privateCmd_HuntSound,
	"/枪声目录":  privateCmd_HuntSoundList,
}
var groupCommandList = map[string]func(cmd_params){
	"":       groupCmd_null, // 空指令，不做任何处理
	"/help":  groupCmd_help,
	"/chat":  groupCmd_chat,
	"/涩图":    groupCmd_setu,
	"/涩图r18": groupCmd_r18,
	"/枪声":    groupCmd_HuntSound,
	"/枪声目录":  groupCmd_HuntSoundList,
	"/禁言抽奖":  groupCmd_BanLottery,
}

// ---------私聊指令处理函数---------
// 空指令
func privateCmd_null(params cmd_params) {
	conn := params.conn
	receivedMsgEvent := params.receivedMsgEvent

	log.Printf("将对私聊回复,msgID:%d,UserID:%d", receivedMsgEvent.MessageID, receivedMsgEvent.UserID)
	// 消息处理
	message_reply := msgHandler(&receivedMsgEvent)
	send_msg(conn, receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, message_reply, false)
}

// help 指令
func privateCmd_help(params cmd_params) {
	conn := params.conn
	receivedMsgEvent := params.receivedMsgEvent

	log.Printf("将对私聊回复,msgID:%d,UserID:%d", receivedMsgEvent.MessageID, receivedMsgEvent.UserID)
	send_msg(conn, receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, "目前支持的指令有：\n/help\n/涩图\n/涩图r18\n/枪声\n/枪声目录", false)
}

// 涩图 指令
func privateCmd_setu(params cmd_params) {
	conn := params.conn
	receivedMsgEvent := params.receivedMsgEvent
	tags := params.tags
	num := params.num
	UserID := receivedMsgEvent.UserID
	GroupID := receivedMsgEvent.GroupID

	log.Println("将对私聊发送涩图 tag:", tags)
	// 得到色图消息
	message_reply := get_setu_MsgReply(tags, 0, num)
	if message_reply == nil {
		send_msg(conn, receivedMsgEvent.MessageType, UserID, GroupID, "涩图获取失败,tag搜索不到图片...", false)
	}
	// 发送消息
	send_private_forward_msg(conn, UserID, message_reply)
}

// r18 指令
func privateCmd_r18(params cmd_params) {
	conn := params.conn
	receivedMsgEvent := params.receivedMsgEvent
	tags := params.tags
	num := params.num
	UserID := receivedMsgEvent.UserID
	GroupID := receivedMsgEvent.GroupID

	log.Println("将对私聊发送r18涩图,tags:", tags)
	// 得到色图消息
	message_reply := get_setu_MsgReply(tags, 1, num)
	if message_reply == nil {
		send_msg(conn, receivedMsgEvent.MessageType, UserID, GroupID, "涩图获取失败,tag搜索不到图片...", false)
	}
	// 发送消息
	send_private_forward_msg(conn, UserID, message_reply)
}

// 枪声 指令
func privateCmd_HuntSound(params cmd_params) {
	conn := params.conn
	receivedMsgEvent := params.receivedMsgEvent

	// 解析cq码，获取无cq格式的消息内容
	cqmsg := ParseCQmsg(receivedMsgEvent.Message)
	log.Println("将对私聊发送枪声:" + GetCQCode_HuntSound(cqmsg.Text))
	send_msg(conn, receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, GetCQCode_HuntSound(cqmsg.Text), false)
}

// 枪声目录 指令
func privateCmd_HuntSoundList(params cmd_params) {
	conn := params.conn
	receivedMsgEvent := params.receivedMsgEvent

	log.Println("将对私聊发送枪声目录")
	send_msg(conn, receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, utils.GetGunIndex(), false)
}

// ---------群聊指令处理函数---------
// null 指令
func groupCmd_null(params cmd_params) {
	log.Printf("非指令消息")
}

// help 指令
func groupCmd_help(params cmd_params) {
	log.Printf("将对群聊回复,msgID:%d,UserID:%d,GroupID:%d", receivedMsgEvent.MessageID, receivedMsgEvent.UserID, receivedMsgEvent.GroupID)
	send_msg(params.conn, receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, "目前支持的指令有：\n/help\n/chat\n/涩图\n/涩图r18\n/枪声\n/枪声目录\n/禁言抽奖", false)
}

// chat 指令
func groupCmd_chat(params cmd_params) {
	conn := params.conn
	receivedMsgEvent := params.receivedMsgEvent

	log.Printf("将对群聊回复,msgID:%d,UserID:%d,GroupID:%d", receivedMsgEvent.MessageID, receivedMsgEvent.UserID, receivedMsgEvent.GroupID)
	// 消息处理
	message_reply := msgHandler(&receivedMsgEvent)
	send_msg(conn, receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, message_reply, false)
}

// 涩图 指令
func groupCmd_setu(params cmd_params) {
	conn := params.conn
	receivedMsgEvent := params.receivedMsgEvent
	tags := params.tags
	num := params.num
	UserID := receivedMsgEvent.UserID
	GroupID := receivedMsgEvent.GroupID

	log.Println("将对群聊发送涩图 tags:", tags)
	// 得到色图消息
	message_reply := get_setu_MsgReply(tags, 0, num)
	if message_reply == nil {
		send_msg(conn, receivedMsgEvent.MessageType, UserID, GroupID, "涩图获取失败,tag搜索不到图片...", false)
	}
	// 发送消息
	send_group_forward_msg(conn, GroupID, message_reply)
}

// r18 指令
func groupCmd_r18(params cmd_params) {
	conn := params.conn
	receivedMsgEvent := params.receivedMsgEvent
	tags := params.tags
	num := params.num
	UserID := receivedMsgEvent.UserID
	GroupID := receivedMsgEvent.GroupID

	log.Println("将对群聊发送r18涩图 tags:", tags)
	// 得到色图消息
	message_reply := get_setu_MsgReply(tags, 1, num)
	if message_reply == nil {
		send_msg(conn, receivedMsgEvent.MessageType, UserID, GroupID, "涩图获取失败,tag搜索不到图片...", false)
	}
	// 发送消息
	send_group_forward_msg(conn, GroupID, message_reply)
}

// 枪声 指令
func groupCmd_HuntSound(params cmd_params) {
	conn := params.conn
	receivedMsgEvent := params.receivedMsgEvent

	// 解析cq码，获取无cq格式的消息内容'
	cqmsg := ParseCQmsg(receivedMsgEvent.Message)
	log.Println("将对群聊发送枪声" + GetCQCode_HuntSound(cqmsg.Text))
	send_msg(conn, receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, GetCQCode_HuntSound(cqmsg.Text), false)
}

// 枪声目录 指令
func groupCmd_HuntSoundList(params cmd_params) {
	conn := params.conn
	receivedMsgEvent := params.receivedMsgEvent

	log.Println("将对群聊发送枪声目录")
	send_msg(conn, receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, utils.GetGunIndex(), false)
}

// 禁言抽奖 指令
func groupCmd_BanLottery(params cmd_params) {
	conn := params.conn
	receivedMsgEvent := params.receivedMsgEvent
	time := rand.Intn(60) + 1
	log.Printf("将对群聊:%d,禁言qq用户:%d,时间:%d", receivedMsgEvent.GroupID, receivedMsgEvent.UserID, time)
	set_group_ban(conn, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, time)
	send_msg(conn, receivedMsgEvent.MessageType, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, "禁言抽奖开始，禁言时间为"+fmt.Sprintf("%d", time)+"秒", false)
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
