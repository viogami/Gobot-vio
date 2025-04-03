package command

import (
	"log/slog"

	"github.com/viogami/Gobot-vio/AI"
	"github.com/viogami/Gobot-vio/gocq"
)

type cmdChat struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdChat) Execute(params CommandParams) {
	sender := gocq.Instance.Sender

	reply := AI.NewAIServer().ProcessMessage(params.Message)
	msgParams := gocq.SendMsgParams{
		MessageType: params.MessageType,
		UserID:      params.UserId,
		GroupID:     params.GroupId,
		Message:     reply,
		AutoEscape:  false,
	}
	slog.Info("调用ai执行指令:/chat", "reply", reply)

	sender.SendMsg(msgParams)
}

func (c *cmdChat) GetInfo(index int) string {
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

func newCmdChat() *cmdChat {
	return &cmdChat{
		Command:     "/chat",
		Description: "聊天指令",
		CmdType:     "all",
	}
}
