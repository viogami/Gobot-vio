package command

import (
	"github.com/viogami/Gobot-vio/AIServer"
	"github.com/viogami/Gobot-vio/gocq"
)

type cmdChat struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdChat) Execute(params CommandParams) {
	reply := AIServer.NewAIServer().ProcessMessage(params.Message)
	
	msgParams := gocq.MsgSendParams{
		MessageType: params.MessageType,
		UserID: 	params.UserId,
		GroupID: 	params.GroupId,
		Message:     reply,
		AutoEscape:  false,
	}
	gocq.MsgSend(msgParams)
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

func NewCmdChat() *cmdChat {
	return &cmdChat{
		Command:     "/chat",
		Description: "聊天指令",
		CmdType:     "group",
	}
}
