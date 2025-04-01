package event

import (
	"encoding/json"
	"log/slog"
	"regexp"

	"github.com/viogami/Gobot-vio/gocq"
	"github.com/viogami/Gobot-vio/gocq/command"
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
	cqmsg := gocq.ParseCQmsg(m.Message)
	f := m.parseCommand(cqmsg)
	if f == nil {
		slog.Warn("MessageEvent", "msg", "接受到非私聊或者非指令的群聊消息", "command", cqmsg.Text)
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


func (m *MessageEvent) parseCommand(cqmsg gocq.CQmsg) command.Command {
	cmdStr := "/chat"
	// 判断是否at我,若否，则将命令为空
	Atme := cqmsg.IsAtme(m.SelfID)
	if !Atme {
		cmdStr = ""
	}
	// 定义正则表达式匹配以中文字符开头的命令
	commandPattern := regexp.MustCompile(`^/([^ ]+)`)
	cmdStr = commandPattern.FindString(cqmsg.Text)

	return command.CommandMap[cmdStr]
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
