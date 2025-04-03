package gocq

type RHttpResq struct {
	Status  string         `json:"status"`
	Retcode int            `json:"retcode"`
	Msg     string         `json:"msg"`
	Wording string         `json:"wording"`
	Data    map[string]any `json:"data"`
	Echo    string         `json:"echo"`
}

type RSender struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	// Sex      string `json:"sex"`
	// Age      int32  `json:"age"`
}

type RSendMsg struct {
	MessageId int32 `json:"message_id"` // 消息id
}

type RGetMsg struct {
	Group       bool    `json:"group"`        // 是否是群消息
	GroupId     int64   `json:"group_id"`     // 群号
	MessageId   int32   `json:"message_id"`   // 消息id
	RealId      int32   `json:"real_id"`      // 消息发送者的真实id
	MessageType string  `json:"message_type"` // 消息类型
	Sender      RSender `json:"sender"`       // 消息发送者
	Time        int32   `json:"time"`         // 消息发送时间
	Message     string  `json:"message"`      // 消息内容
	RawMessage  string  `json:"raw_message"`  // 原始消息内容
}
