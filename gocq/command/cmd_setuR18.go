package command

import "github.com/viogami/Gobot-vio/gocq"

type cmdSetuR18 struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdSetuR18) Execute(params CommandParams) {
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

func (c *cmdSetuR18) GetInfo(index int) string {
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

func NewCmdSetuR18() *cmdSetuR18 {
	return &cmdSetuR18{
		Command:     "/涩图r18",
		Description: "随机r18涩图,规则同上",
		CmdType:     "all",
	}
}
