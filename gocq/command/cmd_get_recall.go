package command

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/viogami/Gobot-vio/gocq"
	"github.com/viogami/Gobot-vio/utils"
)

type redisRecord struct {
	MessageId  int32       `json:"message_id"`
	OperatorId json.Number `json:"operator_id"`
	UserId     json.Number `json:"user_id"`
}

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
	record, err := client.LRange(context.Background(), key, 1, -1).Result()
	if err != nil {
		slog.Error("获取上一次撤回的消息 ID 失败", "error", err)
		return
	}
	redisData := new(redisRecord)
	if err := json.Unmarshal([]byte(record[0]), &redisData); err != nil {
		slog.Error("解析上一次撤回的消息 ID 失败", "error", err)
		return
	}
	// 获取消息 ID 和消息内容
	messageId := redisData.MessageId
	userId := redisData.UserId
	operatorId := redisData.OperatorId

	resp := sender.GetMsg(messageId)
	time := utils.Time2Str(resp["time"])

	reply := fmt.Sprintf("撤回时间:%s\n发送者:%s\n撤回者:%s\n消息内容:%s", time, userId, operatorId, resp["message"])
	msgParams := gocq.SendMsgParams{
		MessageType: params.MessageType,
		GroupID:     params.GroupId,
		UserID:      params.UserId,
		Message:     reply,
		AutoEscape:  false,
	}
	sender.SendMsg(msgParams)
	slog.Info("执行指令:撤回了什么")
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
		CmdType:     COMMAND_TYPE_GROUP,
	}
}
