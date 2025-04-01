package event

import (
	"encoding/json"
	"fmt"
)

type IEvent interface {
	LogInfo()
	Handle()
}

type Event struct {
	PostType string `json:"post_type"`
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
}

// 解析事件 JSON 数据
func ParseEvent(data []byte) (IEvent, error) {
	var t Event
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	// 根据 type 解析不同的事件
	switch t.PostType {
	case "message":
		var e *MessageEvent
		if err := json.Unmarshal(data, &e); err != nil {
			return nil, err
		}
		return e, nil
	case "notice":
		var e *NoticeEvent
		if err := json.Unmarshal(data, &e); err != nil {
			return nil, err
		}
		return e, nil
	case "request":
		var e *RequestEvent
		if err := json.Unmarshal(data, &e); err != nil {
			return nil, err
		}
		return e, nil
	case "meta_event":
		var e *MetaEvent
		if err := json.Unmarshal(data, &e); err != nil {
			return nil, err
		}
		return e, nil
	default:
		return nil, fmt.Errorf("unknown event type: %s", t.PostType)
	}
}
