package command

import (
	"log/slog"

	"github.com/viogami/Gobot-vio/gocq"
)

type cmdHelp struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdHelp) Execute(params CommandParams) {
	reply := ""
	if params.MessageType == "group" {
		reply = c.groupReply()
		slog.Info("群聊指令列表,群号:%d,UserID:%d", params.GroupId, params.UserId)
	} else if params.MessageType == "private" {
		reply = c.privateReply()
		slog.Info("私聊指令列表,msgID:%d,UserID:%d", params.MessageId, params.UserId)
	}
	msgParams := gocq.MsgSendParams{
		MessageType: params.MessageType,
		UserID:      params.UserId,
		GroupID:     params.GroupId,
		Message:     reply,
		AutoEscape:  false,
	}
	gocq.MsgSend(msgParams)
}

func (c *cmdHelp) GetInfo(index int) string {
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

func (c *cmdHelp) privateReply() string {
	reply := "指令列表:\n"
	for _, v := range CommandList {
		if v.GetInfo(2) == "private" || v.GetInfo(2) == "all" {
			reply += v.GetInfo(0) + ":" + v.GetInfo(1) + "\n"
		}
	}
	return reply
}

func (c *cmdHelp) groupReply() string {
	reply := "指令列表:\n"
	for _, v := range CommandList {
		if v.GetInfo(2) == "group" || v.GetInfo(2) == "all" {
			reply += v.GetInfo(0) + ":" + v.GetInfo(1) + "\n"
		}
	}
	return reply

}

func NewCmdHelp() *cmdHelp {
	return &cmdHelp{
		Command:     "/help",
		Description: "查看帮助",
		CmdType:     "all",
	}
}
