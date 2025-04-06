package utils

import "time"

const (
	TIME_FORMAT  = "2006-01-02 15:04:05"
	TIME_FORMAT2 = "2006-01-02 15:04:05.000"
)

func TimeToStr(t any) string {
	if t == nil {
		return ""
	}
	switch t := t.(type) {
	case float64:
		return time.Unix(int64(t), 0).Format(TIME_FORMAT)
	case int64:
		return time.Unix(t, 0).Format(TIME_FORMAT)
	case string:
		return t
	default:
		return ""
	}
}
