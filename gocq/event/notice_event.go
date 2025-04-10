package event

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/viogami/viogo/gocq"
)

var POST_NOTICE_TYPE2STR = map[string]string{
	"group_upload":   "群文件上传",
	"group_admin":    "群管理员变更",
	"group_decrease": "群成员减少",
	"group_increase": "群成员增加",
	"group_ban":      "群成员禁言",
	"friend_add":     "好友添加",
	"group_recall":   "群消息撤回",
	"friend_recall":  "好友消息撤回",
	"group_card":     "群名片变更",
	"offline_file":   "离线文件上传",
	"client_status":  "客户端状态变更",
	"essence":        "精华消息",
	"notify":         "系统通知",
}

var POST_NOTICE_SUB_TYPE2STR = map[string]string{
	"honor":      "群荣誉变更",
	"poke":       "戳一戳",
	"lucky_king": "群红包幸运王",
	"title":      "群成员头衔变更",
}

type NoticeEvent struct {
	Event
	NoticeType string `json:"notice_type"`

	UserID     int64 `json:"user_id"`
	GroupID    int64 `json:"group_id"`
	OperatorId int64 `json:"operator_id"`
	MessageId  int64 `json:"message_id"`
}

func (n *NoticeEvent) LogInfo() {
	slog.Info("NoticeEvent",
		"notice_type", POST_NOTICE_TYPE2STR[n.NoticeType],
		"user_id", n.UserID,
		"group_id", n.GroupID,
		"operator_id", n.OperatorId,
		"message_id", n.MessageId,
	)
}

func (n *NoticeEvent) Handle() {
	notice_type := n.NoticeType
	groupId := n.GroupID
	userId := n.UserID
	opId := n.OperatorId
	msgId := n.MessageId

	sender := gocq.Instance.Sender
	params := gocq.SendMsgParams{
		MessageType: "group",
		GroupID:     groupId,
		UserID:      userId,
		AutoEscape:  false,
	}
	switch notice_type {
	// 群成员增加
	case "group_increase":
		params.Message = "欢迎~需要我就@我~输入'help',查看bot指令列表~"
		sender.SendMsg(params)
	// 群成员减少
	case "group_decrease":
		params.Message = "再见了宝宝~"
		sender.SendMsg(params)
	// 消息撤回
	case "group_recall":
		// 将撤回消息存储为有序列表中的JSON字符串
		recallData, _ := json.Marshal(map[string]any{
			"message_id":  msgId,
			"user_id":     userId,
			"operator_id": opId,
		})
		// 添加到群聊对应的撤回消息列表
		key := fmt.Sprintf("group_recall_%d", groupId)
		pushToRedis(key, string(recallData), 72) // 将消息推送到redis,过期时间72小时
	}
}

func NewNoticeEvent(p []byte) (*NoticeEvent, error) {
	noticeEvent := new(NoticeEvent)
	err := json.Unmarshal(p, &noticeEvent)
	if err != nil {
		return nil, err
	}
	return noticeEvent, nil
}
