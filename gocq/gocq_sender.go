package gocq

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/viogami/Gobot-vio/gocq/cqCode"
)

type GocqSender struct {
	writeMutex sync.Mutex // 添加互斥锁，ws无并发安全
	conn       *websocket.Conn
}

func NewGocqSender(conn *websocket.Conn) *GocqSender {
	return &GocqSender{
		writeMutex: sync.Mutex{}, // 初始化互斥锁
		conn:       conn,
	}
}

func (s *GocqSender) sendToGocq(action string, params map[string]any) (resp map[string]interface{}, err error) {
	s.writeMutex.Lock()
	defer s.writeMutex.Unlock()

	messageSend := map[string]interface{}{
		"action": action,
		"params": params,
	}

	err = s.conn.WriteJSON(messageSend)
	if err != nil {
		return nil, err
	}
	// 等待响应
	var response map[string]interface{}
	err = s.conn.ReadJSON(&response)
	if err != nil {
		return nil, err
	}
	if response["status"] != "ok" {
		return nil, fmt.Errorf("调用gocq api失败: %s", response["message"])
	}
	slog.Info("调用gocq api成功", "action", action, "params", params)
	return nil, nil
}

func (s *GocqSender) SendMsg(params SendMsgParams) {
	action := "send_msg"

	if params.MessageType == "group" {
		cq := cqCode.CQCode{
			Type: "at",
			Data: map[string]interface{}{
				"qq": fmt.Sprintf("%d", params.UserID),
			},
		}
		params.Message = cq.GenerateCQCode() + params.Message
	}
	_, err := s.sendToGocq(action, params.toMap())
	if err != nil {
		slog.Error("发送消息失败", "error", err)
		return
	}
}

func (s *GocqSender) SendGroupForwardMsg(params SendGroupForwardMsgParams) {
	action := "send_group_forward_msg"

	_, err := s.sendToGocq(action, params.toMap())
	if err != nil {
		slog.Error("发送群聊合并消息失败", "error", err)
		return
	}
}

func (s *GocqSender) SendPrivateForwardMsg(params SendPrivateForwardMsgParams) {
	action := "send_private_forward_msg"

	_, err := s.sendToGocq(action, params.toMap())
	if err != nil {
		slog.Error("发送私聊合并消息失败", "error", err)
		return
	}
}

func (s *GocqSender) SetGroupBan(params SendSetGroupBanParams) {
	action := "set_group_ban"

	_, err := s.sendToGocq(action, params.toMap())
	if err != nil {
		slog.Error("设置群禁言失败", "error", err)
		return
	}
}

func (s *GocqSender) GetMsg(msgid int32)map[string]interface{} {
	action := "get_msg"
	params := map[string]interface{}{
		"message_id": msgid,
	}

	resp, err := s.sendToGocq(action, params)
	if err != nil {
		slog.Error("获取消息失败", "error", err)
		return nil
	}
	return resp
}
