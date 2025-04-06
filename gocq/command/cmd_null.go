package command

import (
	"log/slog"
)

type cmdNull struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdNull) Execute(params CommandParams) {
	slog.Info("空指令")
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
		CmdType:     COMMAND_TYPE_ALL,
	}
}
