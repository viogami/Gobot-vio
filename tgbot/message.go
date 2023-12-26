package tgbot

import (
	"Gobot-vio/chatgpt"
	"log"
	"math/rand"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// 检查当前是否应该发送消息,私有
func checksmg(message *tgbotapi.Message) bool {
	var issend bool
	if message.Chat != nil {
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

// 处理用户消息逻辑，公有
func HandleIncomingMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	// 分析消息数据
	uid := message.From.ID
	MessageID := message.MessageID
	gid := message.Chat.ID
	UserName := message.From.UserName
	text := message.Text
	// 是否发送消息触发器
	issend := checksmg(message)

	//定义回复的message
	var msg tgbotapi.MessageConfig

	if issend {
		// 定义回复信息的数组
		msg.Text = "你好,即将调用gpt3.5turbo的API"
		if UserName == "viogami" {
			msg.Text = "主人你好,即将为你调用gpt3.5turbo的API~"
		}
		sendMessage(msg)

		// 调用ChatGPT API
		gptResponse, err := chatgpt.InvokeChatGPTAPI(text)
		if err != nil {
			log.Printf("Error calling ChatGPT API: %v", err)
			gptResponse = "gpt调用失败了😥 错误信息：\n" + err.Error()
		}

		if message.Chat.IsSuperGroup() || message.Chat.IsGroup() {
			msg = tgbotapi.NewMessage(gid, gptResponse)
		} else {
			msg = tgbotapi.NewMessage(uid, gptResponse)
		}
		msg.ReplyToMessageID = MessageID //@发信息的人回复
		_, err = bot.Send(msg)
		if err != nil {
			log.Println("Error sending message to user:", err)
		}
	}

	//机器人命令
	switch message.Command() {
	case "start", "help":
		msg.Text = "我是用go编写的bot:vio,我能够基于chatgpt进行回复,并可以自动回复特定关键词"
		sendMessage(msg)
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
	// 		sendMessage(msg)
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
	// 		sendMessage(msg)
	// 	}
	// case "list":
	// 	if checkAdmin(gid, *message.From) {
	// 		rulelists := getRuleList(gid)
	// 		msg.Text = "ID: " + strconv.FormatInt(gid, 10)
	// 		msg.ParseMode = "Markdown"
	// 		msg.DisableWebPagePreview = true
	// 		sendMessage(msg)
	// 		for _, rlist := range rulelists {
	// 			msg.Text = rlist
	// 			msg.ParseMode = "Markdown"
	// 			msg.DisableWebPagePreview = true
	// 			sendMessage(msg)
	// 		}
	// 	}
	case "admin":
		msg.Text = "[" + message.From.String() + "](tg://user?id=" + strconv.FormatInt(uid, 10) + ") 请求管理员出来打屁股\r\n\r\n" + getAdmins(gid)
		msg.ParseMode = "Markdown"
		sendMessage(msg)
		banMember(gid, uid, 30)
	case "banme":
		botme, _ := bot.GetChatMember(tgbotapi.GetChatMemberConfig{
			ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
				ChatID: gid,
				UserID: uid}})
		if botme.CanRestrictMembers {
			sec := rand.Intn(10) + 5
			banMember(gid, uid, int64(sec))
			msg.Text = "恭喜[" + message.From.String() + "](tg://user?id=" + strconv.FormatInt(uid, 10) + ")获得" + strconv.Itoa(sec) + "秒的禁言礼包"
			msg.ParseMode = "Markdown"
		} else {
			msg.Text = "请给我禁言权限,否则无法进行"
		}
		sendMessage(msg)
	case "me":
		myuser := message.From
		msg.Text = "[" + message.From.String() + "](tg://user?id=" + strconv.FormatInt(uid, 10) + ") 的账号信息" +
			"\r\nID: " + strconv.FormatInt(uid, 10) +
			"\r\nUseName: [" + message.From.String() + "](tg://user?id=" + strconv.FormatInt(uid, 10) + ")" +
			"\r\nLastName: " + myuser.LastName +
			"\r\nFirstName: " + myuser.FirstName +
			"\r\nIsBot: " + strconv.FormatBool(myuser.IsBot)
		msg.ParseMode = "Markdown"
		sendMessage(msg)
	default:
	}
}
