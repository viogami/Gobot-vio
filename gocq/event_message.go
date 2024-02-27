package gocq

import (
	"encoding/json"
	"log"
)

// 下面的Message结构体为array形式，并未使用
// 当前使用的是string形式，修改请到gocq的config文件中改变上报属性
type Message struct {
	// Type string      `json:"type"`
	// Data interface{} `json:"data"`
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

// 私聊消息
type PrivateMessage struct {
	Time        int64  `json:"time"`
	SelfID      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageID   int32  `json:"message_id"`
	UserID      int64  `json:"user_id"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
	Font        int32  `json:"font"`
	Sender      Sender `json:"sender"`
	TargetID    int64  `json:"target_id"`
	TempSource  int    `json:"temp_source"`
}

// 群聊消息
type GroupMessage struct {
	Time        int64     `json:"time"`
	SelfID      int64     `json:"self_id"`
	PostType    string    `json:"post_type"`
	MessageType string    `json:"message_type"`
	SubType     string    `json:"sub_type"`
	MessageID   int32     `json:"message_id"`
	UserID      int64     `json:"user_id"`
	Message     string    `json:"message"`
	RawMessage  string    `json:"raw_message"`
	Font        int32     `json:"font"`
	Sender      Sender    `json:"sender"`
	GroupID     int64     `json:"group_id"`
	Anonymous   Anonymous `json:"anonymous"`
}

type Anonymous struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

func get_msg_info(p []byte, msgType string) interface{} {
	var privateMessage PrivateMessage
	var groupMessage GroupMessage
	switch msgType {
	case "private":
		err := json.Unmarshal(p, &privateMessage)
		if err != nil {
			log.Println("Error parsing JSON to privateMessage:", err)
		}
		return privateMessage
	case "group":
		err := json.Unmarshal(p, &groupMessage)
		if err != nil {
			log.Println("Error parsing JSON to groupMessage:", err)
		}
		return groupMessage
	}
	return nil
}
