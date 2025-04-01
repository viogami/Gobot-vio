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
		return
	}
	reply := AIServer.NewAIServer().ProcessMessage(params.Message)
	msgParams := gocq.SendMsgParams{
		MessageType: params.MessageType,
		UserID:      params.UserId,
		GroupID:     params.GroupId,
		Message:     reply,
		AutoEscape:  false,
	}
	slog.Info("调用ai执行私聊", "reply", reply)
	sender := gocq.NewGocqSender()
	sender.SendMsg(msgParams)
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

func newCmdNull() *cmdNull {
	return &cmdNull{
		Command:     "",
		Description: "空指令",
		CmdType:     "all",
	}
}
