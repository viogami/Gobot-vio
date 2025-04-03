package command

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/viogami/Gobot-vio/gocq"
)

type cmdGetRecall struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdGetRecall) Execute(params CommandParams) {
	client := gocq.Instance.RedisClient
	sender := gocq.Instance.Sender
	// 从 Redis 中获取上一次撤回的消息 ID
	key := fmt.Sprintf("%d", params.GroupId)
	message, err := client.LRange(context.Background(), key, 1, -1).Result()
	if err != nil {
		slog.Error("获取上一次撤回的消息 ID 失败", "error", err)
		return
	}
	var res map[string]any
	if err := json.Unmarshal([]byte(message[0]), &res); err != nil {
		slog.Error("解析上一次撤回的消息 ID 失败", "error", err)
		return
	}
	// 获取消息 ID 和消息内容
	messageId := int32(res["message_id"].(float64))
	operatorId := res["operator_id"]
	userId := res["user_id"]

	resp := sender.GetMsg(messageId)

	msgParams := gocq.SendMsgParams{
		MessageType: params.MessageType,
		GroupID:     params.GroupId,
		UserID:      params.UserId,
		Message:     resp["message"].(string),
		AutoEscape:  false,
	}
	sender.SendMsg(msgParams)

	reply := fmt.Sprintf("时间:%d,发送者:%d,撤回人:%d", resp["time"], userId, operatorId)
	msgParams = gocq.SendMsgParams{
		MessageType: params.MessageType,
		GroupID:     params.GroupId,
		UserID:      params.UserId,
		Message:     resp["message"].(string),
		AutoEscape:  false,
	}
	sender.SendMsg(msgParams)
	slog.Info("执行指令:撤回了什么", "reply", reply)
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
		Command:     "撤回了什么",
		Description: "获取上一条撤回消息",
		CmdType:     "group",
	}
}
