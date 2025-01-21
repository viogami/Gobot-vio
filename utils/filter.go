package utils

var FilterTextMap = map[string]string{
	"viogami": "谢谢你提及我的主人~",
	"习近平":     "请不要提及敏感词汇~",
	"习大大":     "请不要提及敏感词汇~",
}

func Master_ID(userID int64) bool {
	return userID == 123456789
}

func Msg_Filter(text string) string {
	if FilterTextMap[text] != "" {
		return FilterTextMap[text]
	}
	return ""
}

