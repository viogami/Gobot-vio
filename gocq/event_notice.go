package gocq

import (
	"encoding/json"
	"log"
)

type PrivateRecallNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	UserID     int64  `json:"user_id"`
	MessageID  int64  `json:"message_id"`
}

type GroupRecallNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	OperatorID int64  `json:"operator_id"`
	MessageID  int64  `json:"message_id"`
}

type GroupIncreaseNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	OperatorID int64  `json:"operator_id"`
}

type GroupDecreaseNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	OperatorID int64  `json:"operator_id"`
}

func get_notice_info(p []byte, NoticeType string) interface{} {
	var (
		group_recall_notice   GroupRecallNotice
		private_recall_notice PrivateRecallNotice
		group_increase_notice GroupIncreaseNotice
		group_decrease_notice GroupDecreaseNotice
	)
	switch NoticeType {
	case "group_recall":
		err := json.Unmarshal(p, &group_recall_notice)
		if err != nil {
			log.Println("Error parsing JSON to group_recall_notice:", err)
		}
		return group_recall_notice
	case "private_recall":
		err := json.Unmarshal(p, &private_recall_notice)
		if err != nil {
			log.Println("Error parsing JSON to private_recall_notice:", err)
		}
		return private_recall_notice
	case "group_increase":
		err := json.Unmarshal(p, &group_increase_notice)
		if err != nil {
			log.Println("Error parsing JSON to group_increase_notice:", err)
		}
		return group_increase_notice
	case "group_decrease":
		err := json.Unmarshal(p, &group_decrease_notice)
		if err != nil {
			log.Println("Error parsing JSON to group_decrease_notice:", err)
		}
		return group_decrease_notice
	}
	return nil
}
