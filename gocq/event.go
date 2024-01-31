package gocq

import (
	"Gobot-vio/utils"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
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

// 根据事件发送消息
func Send_by_event(conn *websocket.Conn) {
	switch receivedEvent.PostType {
	case "message":
		// 消息事件
		msgtype := receivedMsgEvent.MessageType
		CQcodes := ParseCQmsg(receivedMsgEvent.Message).CQcodes
		msgText := ParseCQmsg(receivedMsgEvent.Message).Text
		// 判断是否at我
		Atme := false
		for _, CQcode := range CQcodes {
			if CQcode.Type == "at" && CQcode.Params["qq"] == fmt.Sprintf("%d", receivedEvent.SelfID) {
				Atme = true
			}
		}
		// 判断是否发送涩图
		Setu, tags := utils.SetuCheck(msgText)

		if msgtype == "private" {
			if Setu {
				log.Println("将对私聊发送涩图 tag:", tags)
				send_image(conn, &receivedMsgEvent, tags, 1, 1)
			} else {
				log.Println("将对私聊回复,receivedMsgEvent:", receivedMsgEvent)
				send_private_msg(conn, &receivedMsgEvent)
			}
		} else if msgtype == "group" && Atme {
			if Setu {
				log.Println("将对私聊发送涩图 tag:", tags)
				send_image(conn, &receivedMsgEvent, tags, 1, 1)
			} else {
				log.Println("将对at我的群聊回复,receivedMsgEvent:", receivedMsgEvent)
				send_group_msg(conn, &receivedMsgEvent)
			}
		} else {
			log.Println("不是私聊或者at我的群聊")
		}
	case "message_sent":
	// 机器人自己发送消息事件
	case "request":
	// 请求事件
	case "notice":
	// 通知事件
	case "meta_event":
		// 元事件
	}
}
