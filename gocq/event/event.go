package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/viogami/viogo/gocq"
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

func IsEvent(data []byte) bool {
	var t Event
	if err := json.Unmarshal(data, &t); err != nil {
		return false
	}
	// 验证必要的事件字段
	if t.PostType == "" || t.Time == 0 || t.SelfID == 0 {
		return false
	}
	return true
}

// ParseEvent 解析事件 JSON 数据
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
		return nil, fmt.Errorf("解析事件失败,raw data: %s", string(data))
	}
}

// pushToRedis 将消息推送到 Redis
func pushToRedis(key string, msg string, expiredTime int) {
	client := gocq.Instance.RedisClient
	if client == nil {
		slog.Error("Redis client is nil")
		return
	}
	// 使用RPUSH添加到群聊对应的消息列表
	err := client.RPush(context.Background(), key, msg).Err()
	if err != nil {
		slog.Error("Failed to push message to redis", "error", err)
		return
	}

	if expiredTime > 0 {
		err = client.Expire(context.Background(), key, time.Duration(expiredTime)*time.Hour).Err()
		if err != nil {
			slog.Error("Failed to set expiration for redis key", "error", err)
			return
		}
	}
	slog.Info("Message pushed to redis", "key", key, "message", msg)
}
