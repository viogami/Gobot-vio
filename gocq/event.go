package gocq

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"regexp"

	"github.com/gorilla/websocket"
	"github.com/viogami/Gobot-vio/utils"
)

type Event struct {
	PostType string `json:"post_type"`
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
}

type MessageEvent struct {
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageID   int32  `json:"message_id"`
	UserID      int64  `json:"user_id"`
	GroupID     int64  `json:"group_id"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
	Font        int    `json:"font"`
	Sender      Sender `json:"sender"`
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
var heart_count = 0

// 判断上报类型
func Log_post_type(p []byte) error {
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
		log.Println("Received ", post_type, ":", receivedMsgEvent.MessageType)
	} else if post_type == "request" {
		// 请求事件
		err := json.Unmarshal(p, &receivedRequestEvent)
		if err != nil {
			log.Println("Error parsing JSON to receivedRequestEvent:", err)
			return err
		}
		log.Println("Received ", post_type, ":", receivedRequestEvent.RequestType)
	} else if post_type == "notice" {
		// 通知事件
		err := json.Unmarshal(p, &receivedNoticeEvent)
		if err != nil {
			log.Println("Error parsing JSON to receivedNoticeEvent:", err)
			return err
		}
		log.Println("Received ", post_type, ":", receivedNoticeEvent.NoticeType)
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
			log.Println("Received ", post_type, ":", receivedMetaEvent.MetaEventType)
		}

		if heart_count == 360 {
			log.Println(receivedMetaEvent.MetaEventType, " 30分钟内OK")
			heart_count = 0
		}
	}
	return nil
}

// 处理上报事件
func Handle_event(conn *websocket.Conn) {
	switch receivedEvent.PostType {
	case "message":
		// 消息事件
		msgtype := receivedMsgEvent.MessageType
		cqmsg := ParseCQmsg(receivedMsgEvent.Message)

		// 定义正则表达式匹配以中文字符开头的命令
		commandPattern := regexp.MustCompile(`^/([^ ]+)`)
		// 使用正则表达式查找匹配的指令
		command := commandPattern.FindString(cqmsg.Text)
		log.Println("command:", command)

		// 判断是否at我
		//Atme := Atme(cqmsg)

		// 涩图tag
		tags := utils.Get_tags(cqmsg.Text)

		if msgtype == "private" {
			switch command {
			case "":
				log.Printf("将对私聊回复,msgID:%d,UserID:%d,msg:%s,raw_msg:%s", receivedMsgEvent.MessageID, receivedMsgEvent.UserID, receivedMsgEvent.Message, receivedMsgEvent.RawMessage)
				// 消息处理
				message_reply := msgHandler(&receivedMsgEvent)
				send_private_msg(conn, &receivedMsgEvent, message_reply)
			case "/help":
				log.Printf("将对私聊回复,msgID:%d,UserID:%d,msg:%s,raw_msg:%s", receivedMsgEvent.MessageID, receivedMsgEvent.UserID, receivedMsgEvent.Message, receivedMsgEvent.RawMessage)
				send_private_msg(conn, &receivedMsgEvent, "目前支持的指令有：\n/help\n/涩图\n/涩图r18")
			case "/涩图":
				log.Println("将对私聊发送涩图 tag:", tags)
				send_private_img(conn, &receivedMsgEvent, tags, 0, 1)
			case "/涩图r18":
				log.Println("将对私聊发送r18涩图 tag:", tags)
				send_private_img(conn, &receivedMsgEvent, tags, 1, 1)
			default:
				log.Printf("将对私聊回复,msgID:%d,UserID:%d,msg:%s,raw_msg:%s", receivedMsgEvent.MessageID, receivedMsgEvent.UserID, receivedMsgEvent.Message, receivedMsgEvent.RawMessage)
				send_private_msg(conn, &receivedMsgEvent, "抱歉，我暂时还无法识别这个指令~")
			}
		} else if msgtype == "group" {
			switch command {
			case "":
				log.Printf("非指令消息,msgID:%d,UserID:%d,GroupID:%d,msg:%s,raw_msg:%s", receivedMsgEvent.MessageID, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, receivedMsgEvent.Message, receivedMsgEvent.RawMessage)
			case "/chat":
				log.Printf("将对群聊回复,msgID:%d,UserID:%d,GroupID:%d,msg:%s,raw_msg:%s", receivedMsgEvent.MessageID, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, receivedMsgEvent.Message, receivedMsgEvent.RawMessage)
				// 消息处理
				message_reply := msgHandler(&receivedMsgEvent)
				send_group_msg(conn, &receivedMsgEvent, message_reply)
			case "/涩图":
				log.Println("将对群聊发送涩图 tags:", tags)
				send_group_img(conn, &receivedMsgEvent, tags, 0, 1)
			case "/涩图r18":
				log.Println("将对群聊发送r18涩图 tags:", tags)
				send_group_img(conn, &receivedMsgEvent, tags, 1, 1)
			case "/禁言抽奖":
				time := rand.Intn(60) + 1
				log.Printf("将对群聊:%d,禁言qq用户:%d,时间:%d", receivedMsgEvent.GroupID, receivedMsgEvent.UserID, time)
				set_group_ban(conn, &receivedMsgEvent, time)
				send_group_msg(conn, &receivedMsgEvent, "已禁言"+fmt.Sprintf("%d", time)+"秒")
			default:
				log.Printf("将对群聊回复,msgID:%d,UserID:%d,GroupID:%d,msg:%s,raw_msg:%s", receivedMsgEvent.MessageID, receivedMsgEvent.UserID, receivedMsgEvent.GroupID, receivedMsgEvent.Message, receivedMsgEvent.RawMessage)
				send_group_msg(conn, &receivedMsgEvent, "抱歉，我暂时还无法识别这个指令~")
			}
		} else {
			log.Println("接受到非私聊或者非指令的群聊消息")
		}

	case "message_sent":
	// 机器人自己发送消息事件

	case "notice":
		log.Println("Received notice:", receivedNoticeEvent)

		notice_type := receivedNoticeEvent.NoticeType
		GroupDecreaseNotice := GroupDecreaseNotice{}

		switch notice_type {
		// 群成员增加
		case "group_increase":
			log.Printf("群成员增加,UserID:%d,GroupID:%d", GroupDecreaseNotice.UserID, GroupDecreaseNotice.GroupID)
			send_group_msg(conn, &receivedMsgEvent, "欢迎新朋友~")
		// 群成员减少
		case "group_decrease":
			log.Printf("群成员减少,UserID:%d,GroupID:%d", receivedMsgEvent.UserID, receivedMsgEvent.GroupID)
		// 消息撤回
		case "group_recall":
			log.Printf("消息撤回,UserID:%d,GroupID:%d", receivedMsgEvent.UserID, receivedMsgEvent.GroupID)
		}

	case "request":
		log.Println("Received request:", receivedRequestEvent)

		request_type := receivedRequestEvent.RequestType
		switch request_type {
		// 使用快速响应
		case "friend":
			fast_resp := map[string]interface{}{
				"approve": true,
				"remark":  "auto approve user",
			}
			err := conn.WriteJSON(fast_resp)
			if err != nil {
				log.Println("Error fast_resp approve:", err)
			}
		case "group":
			fast_resp := map[string]interface{}{
				"approve": false,
				"reason":  "you must notice this to my master:qq2654613995",
			}
			err := conn.WriteJSON(fast_resp)
			if err != nil {
				log.Println("Error fast_resp approve:", err)
			}
		}
	case "meta_event":
		// 元事件
	}
}
