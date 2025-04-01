package command

import (
	"fmt"
	"log/slog"

	"github.com/viogami/Gobot-vio/gocq"
	"github.com/viogami/Gobot-vio/utils"
)

type cmdSetu struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdSetu) Execute(params CommandParams) {
	reply := c.getSetuReply(gocq.SendSetuMsgParams{
		Tags: params.Tags,
		R18:  0,
		Num:  1,
	})
	slog.Info("执行指令:/涩图", "reply", reply)
	sender := gocq.NewGocqSender()

	if params.MessageType == "private" {
		msgParams := gocq.SendPrivateForwardMsgParams{
			UserID:  params.UserId,
			Message: reply,
		}
		sender.SendPrivateForwardMsg(msgParams)
		return
	}
	// 如果是群消息，使用 SendGroupForwardMsg 发送
	msgParams := gocq.SendGroupForwardMsgParams{
		GroupID: params.GroupId,
		Message: reply,
	}
	sender.SendGroupForwardMsg(msgParams)
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

func (c *cmdSetu) getSetuReply(params gocq.SendSetuMsgParams) []gocq.CQCode {
	reply := []gocq.CQCode{
		{
			Type: "node",
			Data: map[string]interface{}{
				"name": "LV",
				"uin":  "1524175162",
				"content": []gocq.CQCode{
					{
						Type: "text",
						Data: map[string]any{
							"text": fmt.Sprintf("tags:%s", params.Tags),
						},
					},
				},
			},
		},
	}
	content := []gocq.CQCode{}
	setuInfo := utils.GetSetu(params.Tags, params.R18, params.Num)
	if setuInfo.Error != "" {
		slog.Error("随机色图api调用出错", "error", setuInfo.Error)
		return nil
	}
	if len(setuInfo.Data) == 0 {
		slog.Error("随机色图api调用出错:tag搜索不到,返回数据为空")
		return nil
	}
	for _, data := range setuInfo.Data {
		content = append(content, gocq.NewCQCode("image", map[string]any{
			"file": data.Urls.Regular,
			"url":  data.Urls.Regular,
		}))
	}
	reply[0].Data["content"] = append(reply[0].Data["content"].([]gocq.CQCode), content...)
	return reply
}

func newCmdSetu() *cmdSetu {
	return &cmdSetu{
		Command:     "/涩图",
		Description: "随机涩图,指令后可接tag,用逗号分隔",
		CmdType:     "all",
	}
}
