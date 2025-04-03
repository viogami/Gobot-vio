package tgbot

import (
	"math/rand"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	chatgpt "github.com/viogami/Gobot-vio/AI/openai"
)

// 检查当前是否应该发送消息
func checksmg(bot *tgbotapi.BotAPI, message *tgbotapi.Message) bool {
	issend := false
	if message.Chat != nil && !message.IsCommand() {
		issend = true
	}
	if message.Chat.IsGroup() && !strings.Contains(message.Text, "@"+bot.Self.UserName) {
		issend = false //普通群组，被@才回复
	}
	if message.Chat.IsSuperGroup() && !strings.Contains(message.Text, "@"+bot.Self.UserName) {
		issend = false //超级群组，被@才回复
	}
	return issend
}

// 处理用户消息逻辑
func HandleIncomingMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	// 创建gpt服务
	gpt := chatgpt.NewChatGPTService()
	// 分析消息数据
	uid := message.From.ID
	gid := message.Chat.ID
	UserName := message.From.UserName
	text := message.Text
	// 是否发送消息
	issend := checksmg(bot, message)

	// 定义回复的message 并初始化
	var replymsg tgbotapi.MessageConfig
	if message.Chat.IsSuperGroup() || message.Chat.IsGroup() {
		replymsg = tgbotapi.NewMessage(message.Chat.ID, "")
	} else {
		replymsg = tgbotapi.NewMessage(message.From.ID, "")
	}

	if issend {
		replymsg.Text = "你好,即将调用gpt3.5turbo的API"
		if UserName == "viogami" {
			replymsg.Text = "主人你好,即将为你调用gpt3.5turbo的API~"
		}
		replymsg.ReplyToMessageID = message.MessageID //@发信息的人回复
		sendMessage(bot, replymsg)

		// 调用ChatGPT API
		gptResponse := gpt.InvokeChatGPTAPI(text)

		replymsg.Text = gptResponse

		replymsg.ReplyToMessageID = message.MessageID //@发信息的人回复
		sendMessage(bot, replymsg)
	}

	//机器人命令
	switch message.Command() {
	case "start", "help":
		replymsg.Text = "我是用go编写的bot:vio,我能够基于chatgpt进行回复,并可以自动回复特定关键词"
		sendMessage(bot, replymsg)
	// case "add":
	// 	if CheckAdmin(gid, *message.From) {
	// 		order := message.CommandArguments()
	// 		if order != "" {
	// 			addRule(gid, order)
	// 			msg.Text = "规则添加成功: " + order
	// 		} else {
	// 			msg.Text = addText
	// 			msg.ParseMode = "Markdown"
	// 			msg.DisableWebPagePreview = true
	// 		}
	// 		SendMessage(msg)
	// 	}
	// case "del":
	// 	if checkAdmin(gid, *message.From) {
	// 		order := message.CommandArguments()
	// 		if order != "" {
	// 			delRule(gid, order)
	// 			msg.Text = "规则删除成功: " + order
	// 		} else {
	// 			msg.Text = delText
	// 			msg.ParseMode = "Markdown"
	// 		}
	// 		SendMessage(msg)
	// 	}
	// case "list":
	// 	if checkAdmin(gid, *message.From) {
	// 		rulelists := getRuleList(gid)
	// 		msg.Text = "ID: " + strconv.FormatInt(gid, 10)
	// 		msg.ParseMode = "Markdown"
	// 		msg.DisableWebPagePreview = true
	// 		SendMessage(msg)
	// 		for _, rlist := range rulelists {
	// 			msg.Text = rlist
	// 			msg.ParseMode = "Markdown"
	// 			msg.DisableWebPagePreview = true
	// 			SendMessage(msg)
	// 		}
	// 	}
	case "admin":
		if message.Chat.IsGroup() || message.Chat.IsSuperGroup() {
			replymsg.Text = "[" + message.From.String() + "](tg://user?id=" + strconv.FormatInt(uid, 10) + ") 请求管理员出来打屁股\r\n\r\n" + getAdmins(bot, gid)
			replymsg.ParseMode = "Markdown"
		} else {
			replymsg.Text = "这是群聊命令，亲~"
		}
		sendMessage(bot, replymsg)
		if !checkAdmin(bot, gid, *message.From) {
			banMember(bot, gid, uid, 30)
		}
	case "banme":
		botme, _ := bot.GetChatMember(tgbotapi.GetChatMemberConfig{
			ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
				ChatID: gid,
				UserID: uid}})
		if botme.CanRestrictMembers {
			sec := rand.Intn(10) + 5
			banMember(bot, gid, uid, int64(sec))
			replymsg.Text = "恭喜[" + message.From.String() + "](tg://user?id=" + strconv.FormatInt(uid, 10) + ")获得" + strconv.Itoa(sec) + "秒的禁言礼包"
			replymsg.ParseMode = "Markdown"
		} else {
			replymsg.Text = "请给我禁言权限,否则无法进行"
		}
		sendMessage(bot, replymsg)
	case "me":
		myuser := message.From
		replymsg.Text = "[" + message.From.String() + "](tg://user?id=" + strconv.FormatInt(uid, 10) + ") 的账号信息" +
			"\r\nID: " + strconv.FormatInt(uid, 10) +
			"\r\nUseName: [" + message.From.String() + "](tg://user?id=" + strconv.FormatInt(uid, 10) + ")" +
			"\r\nLastName: " + myuser.LastName +
			"\r\nFirstName: " + myuser.FirstName +
			"\r\nIsBot: " + strconv.FormatBool(myuser.IsBot)
		replymsg.ParseMode = "Markdown"
		sendMessage(bot, replymsg)
	default:
	}
}
