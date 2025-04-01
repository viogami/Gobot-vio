package command

import (
	"log/slog"

	"github.com/viogami/Gobot-vio/gocq"
)

type cmdSetManager struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdSetManager) Execute(params CommandParams) {
	reply := "coming soon"
	msgParams := gocq.SendMsgParams{
		MessageType: params.MessageType,
		GroupID:     params.GroupId,
		UserID:      params.UserId,
		Message:     reply,
		AutoEscape:  false,
	}
	slog.Info("执行指令:/给我管理", "reply", reply)
	sender := gocq.NewGocqSender()
	sender.SendMsg(msgParams)
}

func (c *cmdSetManager) GetInfo(index int) string {
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

func newCmdSetManager() *cmdSetManager {
	return &cmdSetManager{
		Command:     "/给我管理",
		Description: "设置一个管理给你,目前无效",
		CmdType:     "group",
	}
}
