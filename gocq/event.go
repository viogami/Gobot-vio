package gocq

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Event struct {
	PostType string `json:"post_type"`
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
}

type MessageEvent struct {
	MessageType string  `json:"message_type"`
	SubType     string  `json:"sub_type"`
	MessageID   int32   `json:"message_id"`
	UserID      int64   `json:"user_id"`
	Message     Message `json:"message"`
	RawMessage  string  `json:"raw_message"`
	Font        int     `json:"font"`
	Sender      Sender  `json:"sender"`
}
type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
type AtData struct {
	QQ int64 `json:"qq"`
}
type TextData struct {
	Text string `json:"text"`
}

type Sender struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
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

type P struct {
	P interface{} `json:"p"`
}

// 接收的事件
var (
	receivedEvent        Event
	receivedMsgEvent     MessageEvent
	receivedRequestEvent RequestEvent
	receivedNoticeEvent  NoticeEvent
	receivedMetaEvent    MetaEvent
	pp                   P
)

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
		log.Println(json.Unmarshal(p, &pp))
		err := json.Unmarshal(p, &receivedMsgEvent)
		if err != nil {
			log.Println("Error parsing JSON to receivedMsgEvent:", err)
			return err
		}
		log.Println("Received message_event:", receivedMsgEvent.MessageType)
	} else if post_type == "request" {
		// 请求事件
		err := json.Unmarshal(p, &receivedRequestEvent)
		if err != nil {
			log.Println("Error parsing JSON to receivedRequestEvent:", err)
			return err
		}
		log.Println("Received request_event:", receivedRequestEvent.RequestType)
	} else if post_type == "notice" {
		// 通知事件
		err := json.Unmarshal(p, &receivedNoticeEvent)
		if err != nil {
			log.Println("Error parsing JSON to receivedNoticeEvent:", err)
			return err
		}
		log.Println("Received notice_event:", receivedNoticeEvent.NoticeType)
	} else if post_type == "meta_event" {
		// 元事件
		err := json.Unmarshal(p, &receivedMetaEvent)
		if err != nil {
			log.Println("Error parsing JSON to receivedMetaEvent:", err)
			return err
		}
		log.Println("Received meta_event:", receivedMetaEvent.MetaEventType)
	}
	return nil
}

// 发送消息
func Send_by_event(conn *websocket.Conn) {
	if receivedEvent.PostType == "message" || receivedEvent.PostType == "message_sent" {
		// 消息事件
		msgtype := receivedMsgEvent.SubType
		targetID := receivedMsgEvent.UserID
		message := receivedMsgEvent.Message[0].Data.(string)
		if msgtype == "private" {
			Send_msg(conn, msgtype, targetID, message)
		} else if msgtype == "group" && receivedMsgEvent.Message[0].Type == "at" && receivedMsgEvent.Message[0].Data.(int64) == receivedEvent.SelfID {
			Send_msg(conn, msgtype, targetID, message)
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
