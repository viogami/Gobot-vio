package gocq

import (
	"fmt"
	"log/slog"
)

type GocqSender struct {
}

func NewGocqSender() *GocqSender {
	return &GocqSender{}
}

func (s *GocqSender) SendMsg(params SendMsgParams) {
	action := "send_msg"

	if params.MessageType == "group" {
		cq := CQCode{
			Type: "at",
			Data: map[string]interface{}{
				"qq": fmt.Sprintf("%d", params.UserID),
			},
		}
		params.Message = cq.GenerateCQCode() + params.Message
	}

	err := GocqInstance.SendToGocq(action, params.toMap())
	if err != nil {
		slog.Error("发送消息失败", "error", err)
		return
	}
}

func (s *GocqSender) SendGroupForwardMsg(params SendGroupForwardMsgParams) {
	action := "send_group_forward_msg"

	err := GocqInstance.SendToGocq(action, params.toMap())
	if err != nil {
		slog.Error("发送群聊合并消息失败", "error", err)
		return
	}
}

func (s *GocqSender) SendPrivateForwardMsg(params SendPrivateForwardMsgParams) {
	action := "send_private_forward_msg"

	err := GocqInstance.SendToGocq(action, params.toMap())
	if err != nil {
		slog.Error("发送私聊合并消息失败", "error", err)
		return
	}
}

func (s *GocqSender) SetGroupBan(params SendSetGroupBanParams) {
	action := "set_group_ban"

	err := GocqInstance.SendToGocq(action, params.toMap())
	if err != nil {
		slog.Error("设置群禁言失败", "error", err)
		return
	}
}
