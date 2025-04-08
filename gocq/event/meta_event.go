package event

import (
	"encoding/json"
	"log/slog"
	"time"
)

type MetaEvent struct {
	Event

	MetaEventType     string `json:"meta_event_type"`
	lastHeartbeatTime time.Time // 上次心跳时间
}

func (m *MetaEvent) LogInfo() {
	if m.MetaEventType == "heartbeat" {
		// 获取当前时间
		now := time.Now()

		// 检查是否超过 10 秒没有心跳
		if now.Sub(m.lastHeartbeatTime) > 10*time.Second {
			slog.Warn("超过10秒没有收到心跳事件", "last_heartbeat_time", m.lastHeartbeatTime)
		}
		return
	}
	// 其他元事件类型
	slog.Info("MetaEvent", "meta_event_type", m.MetaEventType)
}

func (m *MetaEvent) Handle() {
	// 处理元事件
	metaEventType := m.MetaEventType
	switch metaEventType {
	case "heartbeat":
		// 心跳事件
		m.lastHeartbeatTime = time.Now()
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
