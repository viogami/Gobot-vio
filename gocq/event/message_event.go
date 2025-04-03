package event

import (
	"encoding/json"
	"log/slog"

	"github.com/viogami/Gobot-vio/gocq/command"
	"github.com/viogami/Gobot-vio/gocq/cqCode"
	"github.com/viogami/Gobot-vio/utils"
)

type MessageEvent struct {
	Event

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

type Sender struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}

func (m *MessageEvent) LogInfo() {
	slog.Info("MessageEvent",
		"message_type", m.MessageType,
		"sub_type", m.SubType,
		"message_id", m.MessageID,
		"user_id", m.UserID,
		"group_id", m.GroupID,
		"message", m.Message,
		// "raw_message", m.RawMessage,
		// "font", m.Font
		"sender_id", m.Sender.UserID,
	)
}

func (m *MessageEvent) Handle() {
	cqmsg := cqCode.ParseCQmsg(m.Message)
	f := m.parseCommand(cqmsg)
	if f == nil {
		slog.Info("MessageEvent", "接受到普通群聊消息", m.Message)
		return
	}
	params := command.CommandParams{
		MessageId:   m.MessageID,
		MessageType: m.MessageType,
		Message:     m.Message,
		GroupId:     m.GroupID,
		UserId:      m.UserID,

		SetuParams: command.SetuParams{
			Tags: utils.ReadTags(cqmsg.Text),
		},
	}

	f.Execute(params)
}

func (m *MessageEvent) parseCommand(cqmsg cqCode.CQmsg) command.Command {
	cmdStr := cqmsg.Text
	r := command.CommandMap[cmdStr]
	if r == nil {
		return command.CommandMap["/chat"]
	}
	// 判断是否是私聊消息
	if m.MessageType == "private" {
		if r.GetInfo(2) == "private" || r.GetInfo(2) == "all" {
			return r
		}
	}
	// 判断是否是群聊消息
	slog.Debug("cqmsg", cmdStr, "selfId", m.SelfID)

	if m.MessageType == "group" && cqmsg.IsAtme(m.SelfID) {
		if r.GetInfo(2) == "group" || r.GetInfo(2) == "all" {
			return r
		}
	}
	// 正则表达式匹配是否是命令格式的消息
	// commandPattern := regexp.MustCompile(`^/([^ ]+)`)
	// cmdStr = commandPattern.FindString(cqmsg.Text)
	return nil
}

func NewMessageEvent(p []byte) (*MessageEvent, error) {
	messageEvent := new(MessageEvent)
	err := json.Unmarshal(p, &messageEvent)
	if err != nil {
		slog.Error("Error parsing JSON to MessageEvent:", "err", err)
		return nil, err
	}
	return messageEvent, nil
}
