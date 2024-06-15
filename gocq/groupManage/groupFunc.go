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

// 设置群管理
// @params
// group_id	int64	-	群号
// user_id	int64	-	要设置管理员的 QQ 号
// enable	boolean	true	true 为设置, false 为取消
func Set_group_manager(groupid,userid int64,enable bool)map[string]interface{} {
		sendMessage := map[string]interface{}{
			"action": "set_group_admin",
			"params": map[string]interface{}{
				"group_id": groupid,
				"user_id":  userid,
				"enable": enable,
			},
			"echo": "echo_test",
		}
		return sendMessage
}
