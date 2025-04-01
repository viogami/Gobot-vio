package event

import (
	"encoding/json"
	"log/slog"
)

type RequestEvent struct {
	Event

	RequestType string `json:"request_type"`
}

func (n *RequestEvent) Slog() {
	slog.Info("RequestEvent", "request_type", n.RequestType)
}

func (n *RequestEvent) Do() {
	// 处理请求事件
	request_type := n.RequestType
	switch request_type {
	case "friend":
		// 好友请求
	case "group":
		// 群组请求
	case "group_invite":
		// 群邀请请求
	default:
		slog.Warn("Unknown request type:", "request_type", request_type)
	}
}

func NewRequestEvent(p []byte) (*RequestEvent, error) {
	requestEvent := new(RequestEvent)
	err := json.Unmarshal(p, &requestEvent)
	if err != nil {
		return nil, err
	}
	return requestEvent, nil
}
