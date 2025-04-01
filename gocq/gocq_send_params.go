package gocq

type SendMsgParams struct {
	MessageType string `json:"message_type"` // 消息类型, 支持 private、group , 分别对应私聊、群组, 如不传入, 则根据传入的 *_id 参数判断
	UserID      int64  `json:"user_id"`      // 对方 QQ 号 ( 消息类型为 private 时需要 )
	GroupID     int64  `json:"group_id"`     // 群号 ( 消息类型为 group 时需要 )
	Message     string `json:"message"`      // 要发送的内容
	AutoEscape  bool   `json:"auto_escape"`  // 消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 message 字段是字符串时有效
}

func (params SendMsgParams) toMap() map[string]interface{} {
	return map[string]interface{}{
		"message_type": params.MessageType,
		"user_id":      params.UserID,
		"group_id":     params.GroupID,
		"message":      params.Message,
		"auto_escape":  params.AutoEscape,
	}
}

type SendGroupForwardMsgParams struct {
	GroupID int64    `json:"group_id"` // 群号
	Message []CQCode `json:"messages"` // 消息列表
}

type SendPrivateForwardMsgParams struct {
	UserID  int64    `json:"user_id"`  // 对方 QQ 号
	Message []CQCode `json:"messages"` // 消息列表
}

type SendSetuMsgParams struct {
	Tags []string `json:"tags"` // 色图标签
	R18  int      `json:"r18"`  // 是否 R18
	Num  int      `json:"num"`  // 色图数量
}

func (params SendSetuMsgParams) toMap() map[string]interface{} {
	return map[string]interface{}{
		"tags": params.Tags,
		"r18":  params.R18,
		"num":  params.Num,
	}
}
