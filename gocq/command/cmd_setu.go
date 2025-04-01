package command

import "github.com/viogami/Gobot-vio/gocq"

type cmdSetu struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdSetu) Execute(params CommandParams) {
	reply := "coming soon"
	msgParams := gocq.MsgSendParams{
		MessageType: params.MessageType,
		GroupID:     params.GroupId,
		UserID:      params.UserId,
		Message:     reply,
		AutoEscape:  false,
	}
	gocq.MsgSend(msgParams)
}

func (c *cmdSetu) GetInfo(index int) string {
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

func NewCmdSetu() *cmdSetu {
	return &cmdSetu{
		Command:     "/涩图",
		Description: "随机涩图，指令后可接tag，用逗号分隔",
		CmdType:     "all",
	}
}
