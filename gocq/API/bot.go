// 有关 Bot 账号的相关 API

package api

// 获取登录号
// @response
// user_id	int64	QQ 号
// nickname	string	QQ 昵称
func Get_login_info() message {
	return message{
		Action: "get_login_info",
		Params: map[string]interface{}{},
		Echo:   "echo_test",
	}
}

//设置登录号资料
//@response 无
func Set_qq_profile(nickname, company, email, collage, personal_note string) message {
	return message{
		Action: "set_qq_profile",
		Params: map[string]interface{}{
			"nickname":      nickname,
			"company":       company,
			"email":         email,
			"collage":       collage,
			"personal_note": personal_note,
		},
		Echo: "echo_test",
	}
}

// 获取企业账号信息,只有企点账号可以调用
// @response 无
func Qidian_get_account_info() message {
	return message{
		Action: "qidian_get_account_info",
		Params: map[string]interface{}{},
		Echo:   "echo_test",
	}
}

//获取在线机型
//@response json数组，foreach:
// model_show	string
// need_pay	boolean
func Get_model_show() message {
	return message{
		Action: "_get_model_show",
		Params: map[string]interface{}{},
		Echo:   "echo_test",
	}
}

//设置在线机型
//@response 无
func Set_model_show(model, model_show string) message {
	return message{
		Action: "_set_model_show",
		Params: map[string]interface{}{
			"model":      model,
			"model_show": model_show,
		},
		Echo: "echo_test",
	}
}

// 获取当前账号的在线客户端列表
// @response clients	Device[]	在线客户端列表
// Device
// app_id	int64	客户端ID
// device_name	string	设备名称
// device_kind	string	设备类型
func Get_online_clients(no_cache bool) message {
	return message{
		Action: "get_online_clients",
		Params: map[string]interface{}{
			"no_cache": no_cache,
		},
		Echo: "echo_test",
	}
}
