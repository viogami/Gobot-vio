package gocq

import (
	"encoding/json"
	"log"
	"regexp"

	"github.com/viogami/Gobot-vio/gocq/cqEvent"
	"github.com/viogami/Gobot-vio/utils"
)

type Event struct {
	PostType string `json:"post_type"`
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
}

type MessageEvent struct {
	MessageType string         `json:"message_type"`
	SubType     string         `json:"sub_type"`
	MessageID   int32          `json:"message_id"`
	UserID      int64          `json:"user_id"`
	GroupID     int64          `json:"group_id"`
	Message     string         `json:"message"`
	RawMessage  string         `json:"raw_message"`
	Font        int            `json:"font"`
	Sender      cqEvent.Sender `json:"sender"`
}

type RequestEvent struct {
	RequestType string `json:"request_type"`
}

type NoticeEvent struct {
	NoticeType string `json:"notice_type"`
}

type MetaEvent struct {
	MetaEventType string `json:"meta_event_type"`
}

// 接收的事件
var (
	receivedEvent        Event
	receivedMsgEvent     MessageEvent
	receivedRequestEvent RequestEvent
	receivedNoticeEvent  NoticeEvent
	receivedMetaEvent    MetaEvent
)

// 回复的消息内容
var sendToGocqList = make([]map[string]interface{}, 0)
// 心跳计数
var heart_count = 0
// 撤回消息ID
var recall_msg_id int64

// 判断上报类型
func Log_post_type(p []byte) error {
	// 接受事件初始化
	receivedEvent = Event{}
	receivedMsgEvent = MessageEvent{}
	receivedRequestEvent = RequestEvent{}
	receivedNoticeEvent = NoticeEvent{}
	receivedMetaEvent = MetaEvent{}

	err := json.Unmarshal(p, &receivedEvent)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return err
	}
	post_type := receivedEvent.PostType

	if post_type == "message" || post_type == "message_sent" {
		// 消息事件
		err := json.Unmarshal(p, &receivedMsgEvent)
		if err != nil {
			log.Println("Error parsing JSON to receivedMsgEvent:", err)
			return err
		}
		log.Printf("Received-->%s:%s,sender:%d,groupid:%d", post_type, receivedMsgEvent.MessageType, receivedMsgEvent.Sender.UserID, receivedMsgEvent.GroupID)
	} else if post_type == "request" {
		// 请求事件
		err := json.Unmarshal(p, &receivedRequestEvent)
		if err != nil {
			log.Println("Error parsing JSON to receivedRequestEvent:", err)
			return err
		}
		log.Printf("Received-->%s:%s", post_type, receivedRequestEvent.RequestType)
	} else if post_type == "notice" {
		// 通知事件
		err := json.Unmarshal(p, &receivedNoticeEvent)
		if err != nil {
			log.Println("Error parsing JSON to receivedNoticeEvent:", err)
			return err
		}
		log.Printf("Received-->%s:%s", post_type, receivedNoticeEvent.NoticeType)
	} else if post_type == "meta_event" {
		// 元事件
		err := json.Unmarshal(p, &receivedMetaEvent)
		if err != nil {
			log.Println("Error parsing JSON to receivedMetaEvent:", err)
			return err
		}

		if receivedMetaEvent.MetaEventType == "heartbeat" {
			heart_count++
		} else {
			log.Printf("Received-->%s:%s", post_type, receivedMetaEvent.MetaEventType)
		}

		if heart_count == 360 {
			log.Println(receivedMetaEvent.MetaEventType, " 30分钟内OK")
			heart_count = 0
		}
	}
	return nil
}

// 处理上报事件
func Handle_event(p []byte) []map[string]interface{} {
	sendToGocqList = make([]map[string]interface{}, 0)
	switch receivedEvent.PostType {
	case "message":
		// 消息事件
		msgtype := receivedMsgEvent.MessageType
		command := ""
		// 解析cq码，获取无cq格式的消息内容
		cqmsg := ParseCQmsg(receivedMsgEvent.Message)
		// 判断是否at我
		Atme := Atme(cqmsg)
		if Atme {
			// 如果at我,则将命令设置为/chat
			command = "/chat"
		} else {
			// 定义正则表达式匹配以中文字符开头的命令
			commandPattern := regexp.MustCompile(`^/([^ ]+)`)
			// 使用正则表达式查找匹配的指令
			command = commandPattern.FindString(cqmsg.Text)
		}
		log.Println("command:", command)

		// 构造命令参数
		params := cmd_params{
			receivedMsgEvent: &receivedMsgEvent,
			tags:             utils.Get_tags(cqmsg.Text),
			num:              1,
		}

		if msgtype == "private" {
			cmd := privateCommandList[command]
			if cmd != nil {
				sendToGocqList = append(sendToGocqList, cmd(params))
				return sendToGocqList
			} else {
				log.Printf("识别到未定义指令,command:%s", command)
			}
		} else if msgtype == "group" {
			cmd := groupCommandList[command]
			if cmd != nil {
				sendToGocqList = append(sendToGocqList, cmd(params))
				return sendToGocqList
			} else {
				log.Printf("识别到未定义指令,command:%s", command)
			}
		} else {
			log.Println("接受到非私聊或者非指令的群聊消息")
		}

	case "message_sent":
	// 机器人自己发送消息事件

	case "notice":
		notice_type := receivedNoticeEvent.NoticeType
		switch notice_type {
		// 群成员增加
		case "group_increase":
			group_increase_info := cqEvent.Get_notice_info(p, receivedNoticeEvent.NoticeType).(cqEvent.GroupIncreaseNotice)
			log.Printf("群成员增加,UserID:%d,GroupID:%d", group_increase_info.UserID, group_increase_info.GroupID)

			sendToGocqList = append(sendToGocqList, msg_send("group", group_increase_info.UserID, group_increase_info.GroupID, "欢迎加入,输入'/help',查看bot指令列表~", false))
			return sendToGocqList
		// 群成员减少
		case "group_decrease":
			group_decrease_info := cqEvent.Get_notice_info(p, receivedNoticeEvent.NoticeType).(cqEvent.GroupDecreaseNotice)
			log.Printf("群成员减少,UserID:%d,GroupID:%d", group_decrease_info.UserID, group_decrease_info.GroupID)

			sendToGocqList = append(sendToGocqList, msg_send("group", group_decrease_info.UserID, group_decrease_info.GroupID, "有人离开了群聊~", false))
			return sendToGocqList
		// 消息撤回
		case "group_recall":
			group_recall_info := cqEvent.Get_notice_info(p, receivedNoticeEvent.NoticeType).(cqEvent.GroupRecallNotice)
			log.Printf("消息撤回,UserID:%d,GroupID:%d", group_recall_info.UserID, group_recall_info.GroupID)
			recall_msg_id = group_recall_info.MessageID
		}

	case "request":
		request_type := receivedRequestEvent.RequestType

		switch request_type {
		// 使用快速响应
		case "friend":
			friend_info := cqEvent.Get_request_info(p, receivedRequestEvent.RequestType).(cqEvent.AddFriendRequest)
			log.Println("好友请求:", friend_info.UserID, friend_info.Comment, friend_info.Flag)

			return sendToGocqList
		case "group":
			group_info := cqEvent.Get_request_info(p, receivedRequestEvent.RequestType).(cqEvent.AddGroupRequest)
			log.Println("群请求:", group_info.GroupID, group_info.UserID, group_info.Comment, group_info.Flag)

			return sendToGocqList
		}
	case "meta_event":
		// 元事件
	}
	return nil
}
