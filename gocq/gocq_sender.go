package gocq

import (
	"fmt"
	"log/slog"
	"net"
	"sync"
	"time"

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

func (s *GocqSender) sendToGocq(action string, params map[string]any) (resp map[string]any, err error) {
	s.writeMutex.Lock()

	// 生成唯一echo值
	echoValue := fmt.Sprintf("%s:%d", action, time.Now().UnixNano())

	// 创建消息
	messageSend := map[string]any{
		"action": action,
		"params": params,
		"echo":   echoValue, // 添加echo字段
	}

	// 发送请求
	err = s.conn.WriteJSON(messageSend)
	s.writeMutex.Unlock() // 发送后立即释放锁，允许其他请求发送

	if err != nil {
		return nil, err
	}

	// 设置超时时间
	deadline := time.Now().Add(5 * time.Second)

	// 循环等待匹配echo的响应
	for time.Now().Before(deadline) {
		// 设置读取超时
		err = s.conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		if err != nil {
			return nil, err
		}

		var r map[string]any
		err = s.conn.ReadJSON(&r)

		// 读取超时，继续循环
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			continue
		}

		// 其他错误
		if err != nil {
			_ = s.conn.SetReadDeadline(time.Time{}) // 重置超时
			return nil, err
		}

		// 检查是否为我们的响应
		if echo, ok := r["echo"].(string); ok && echo == echoValue {
			_ = s.conn.SetReadDeadline(time.Time{}) // 重置超时
			slog.Info("调用gocq api成功", "action", action, "params", params)
			return r, nil
		}
	}

	_ = s.conn.SetReadDeadline(time.Time{}) // 重置超时
	return nil, fmt.Errorf("等待响应超时")
}

func (s *GocqSender) SendMsg(params SendMsgParams) {
	action := "send_msg"

	if params.MessageType == "group" {
		cq := cqCode.CQCode{
			Type: "at",
			Data: map[string]any{
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

func (s *GocqSender) GetMsg(msgid int32) map[string]any {
	action := "get_msg"
	params := map[string]any{
		"message_id": msgid,
	}

	resp, err := s.sendToGocq(action, params)
	if err != nil {
		slog.Error("获取消息失败", "error", err)
		return nil
	}
	return resp
}
