package gocq

import (
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

// 判断上报类型
func Log_post_type(p []byte) error {
	err := json.Unmarshal(p, &receivedEvent)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return err
	}
	post_type := receivedEvent.PostType
	/////////////////////////////
	log.Println(receivedEvent)
	if post_type == "message " || post_type == "message_sent" {
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
		log.Println("Received", post_type, ":", receivedMetaEvent.MetaEventType)
	}
	return nil
}

// 发送消息
func Send_by_event(conn *websocket.Conn) {
	if receivedEvent.PostType == "message" {
		// 消息事件
		msgtype := receivedMsgEvent.MessageType
		CQcodes := ParseCQmsg(receivedMsgEvent.Message).CQcodes
		msgText := ParseCQmsg(receivedMsgEvent.Message).Text
		Atme := false
		// 判断是否at我
		for _, CQcode := range CQcodes {
			if CQcode.Type == "at" && CQcode.Params["qq"] == fmt.Sprintf("%d", receivedEvent.SelfID) {
				Atme = true
			}
		}

		if msgtype == "private" {
			targetID := receivedMsgEvent.UserID
			Send_msg(conn, msgtype, targetID, msgText)
		} else if msgtype == "group" && Atme {
			targetID := receivedMsgEvent.GroupID
			Send_msg(conn, msgtype, targetID, msgText)
		} else {
			log.Println("不是私聊或者at我的群聊")
		}

	} else if receivedEvent.PostType == "request" {
		// 请求事件
	} else if receivedEvent.PostType == "notice" {
		// 通知事件
	} else if receivedEvent.PostType == "meta_event" {
		// 元事件
	}
}
