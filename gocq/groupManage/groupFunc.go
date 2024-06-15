package groupmanage

// 设置群组禁言
func Set_group_ban(UserID int64, GroupID int64, time int) map[string]interface{} {
	// 构建消息结构
	sendMessage := map[string]interface{}{
		"action": "set_group_ban",
		"params": map[string]interface{}{
			"group_id": GroupID,
			"user_id":  UserID,
			"duration": time, // 单位秒，0 表示解除禁言
		},
		"echo": "echo_test",
	}
	return sendMessage
}
