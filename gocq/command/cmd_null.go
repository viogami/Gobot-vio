package command

import (
	"log/slog"

	"github.com/viogami/Gobot-vio/AIServer"
	"github.com/viogami/Gobot-vio/gocq"
)

type cmdNull struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdNull) Execute(params CommandParams) {
	if params.MessageType == "group" {
		slog.Info("群聊空指令,群号:%d,UserID:%d", params.GroupId, params.UserId)
		return
	}
	if params.MessageType == "private" {
		slog.Info("将对私聊回复,msgID:%d,UserID:%d", params.MessageId, params.UserId)
	}
	reply := AIServer.NewAIServer().ProcessMessage(params.Message)
	msgParams := gocq.MsgSendParams{
		MessageType: params.MessageType,
		UserID:      params.UserId,
		GroupID:     params.GroupId,
		Message:     reply,
		AutoEscape:  false,
	}
	gocq.MsgSend(msgParams)
}

func (c *cmdNull) GetInfo(index int) string {
	switch index {
	case 0:
		return c.Command
	case 1:
		return c.Description
	case 2:
		return c.CmdType
	}
	return ""
}

func NewCmdNull() *cmdNull {
	return &cmdNull{
		Command:     "",
		Description: "空指令",
		CmdType:     "all",
	}
}
