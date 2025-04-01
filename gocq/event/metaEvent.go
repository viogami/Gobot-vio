package event

import (
	"encoding/json"
	"log/slog"
)

type MetaEvent struct {
	Event

	MetaEventType string `json:"meta_event_type"`
}

var heartCount int = 0

func (m *MetaEvent) LogInfo() {
	heartCount++
	if heartCount == 50 {
		slog.Info("MetaEvent", "meta_event_type", m.MetaEventType, "heartCount", heartCount)
		heartCount = 0
	}
}

func (m *MetaEvent) Do() {
	// 处理元事件
	metaEventType := m.MetaEventType
	switch metaEventType {
	case "heartbeat":
		// 心跳事件
	case "lifecycle":
		// 生命周期事件
	default:
		slog.Warn("Unknown meta event type:", "meta_event_type", metaEventType)
	}
}

func NewMetaEvent(p []byte) (*MetaEvent, error) {
	metaEvent := new(MetaEvent)
	err := json.Unmarshal(p, &metaEvent)
	if err != nil {
		return nil, err
	}
	return metaEvent, nil
}
