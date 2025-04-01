package command

import "github.com/viogami/Gobot-vio/gocq"

type cmdHuntSound struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdHuntSound) Execute(params CommandParams) {
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

func (c *cmdHuntSound) GetInfo(index int) string {
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

func NewCmdHuntSound() *cmdHuntSound {
	return &cmdHuntSound{
		Command:     "/枪声",
		Description: "随机枪声",
		CmdType:     "all",
	}
}
