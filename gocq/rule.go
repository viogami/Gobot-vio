package gocq

import "strings"

func Master_ID(userID int64) bool {
	return userID == 123456789
}

func Msg_Filter(text string) string {
	if strings.Contains(text, "viogami") {
		return "谢谢你提及我的主人~"
	}
	if strings.Contains(text, "习近平") {
		return "已过滤"
	}
	return ""
}
