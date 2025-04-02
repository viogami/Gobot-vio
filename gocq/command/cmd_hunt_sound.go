package command

import (
	"log/slog"

	"github.com/viogami/Gobot-vio/gocq"
	"github.com/viogami/Gobot-vio/utils"
)

type cmdHuntSound struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdHuntSound) Execute(params CommandParams) {
	hs := utils.NewRandHuntSound() // 随机枪声，固定5m
	reply := c.cqReply(hs.Sound)
	msgParams := gocq.SendMsgParams{
		MessageType: params.MessageType,
		GroupID:     params.GroupId,
		UserID:      params.UserId,
		Message:     reply,
		AutoEscape:  false,
	}
	slog.Info("执行指令:/hunt_sound", "reply", reply)
	sender := gocq.NewGocqSender()
	sender.SendMsg(msgParams)
}

func (c *cmdHuntSound) cqReply(soundUrl string) string {
	ret := gocq.CQCode{
		Type: "record",
		Data: map[string]interface{}{
			"file": soundUrl,
		},
	}
	return ret.GenerateCQCode()
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

func newCmdHuntSound() *cmdHuntSound {
	return &cmdHuntSound{
		Command:     "/枪声",
		Description: "随机枪声",
		CmdType:     "all",
	}
}
