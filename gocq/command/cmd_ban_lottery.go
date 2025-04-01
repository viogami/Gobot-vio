package command

import (
	"log/slog"

	"github.com/viogami/Gobot-vio/gocq"
)

type cmdBanLottery struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdBanLottery) Execute(params CommandParams) {
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

func (c *cmdBanLottery) GetInfo(index int) string {
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

func newCmdBanLottery() *cmdBanLottery {
	return &cmdBanLottery{
		Command:     "/禁言抽奖",
		Description: "禁言抽奖0~180秒",
		CmdType:     "group",
	}
}
