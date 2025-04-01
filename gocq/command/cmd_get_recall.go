package command

import (
	"log/slog"

	"github.com/viogami/Gobot-vio/gocq"
)

type cmdGetRecall struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdGetRecall) Execute(params CommandParams) {
	reply := "coming soon"
	msgParams := gocq.SendMsgParams{
		MessageType: params.MessageType,
		GroupID:     params.GroupId,
		UserID:      params.UserId,
		Message:     reply,
		AutoEscape:  false,
	}
	slog.Info("执行指令:/禁言抽奖", "reply", reply)
	sender := gocq.NewGocqSender()
	sender.SendMsg(msgParams)
}

func (c *cmdGetRecall) GetInfo(index int) string {
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

func newCmdGetRecall() *cmdGetRecall {
	return &cmdGetRecall{
		Command:     "/撤回了什么",
		Description: "获取上一条撤回消息",
		CmdType:     "group",
	}
}
