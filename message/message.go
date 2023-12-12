package message

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleIncomingMessage 处理用户消息
func HandleIncomingMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	// 分析消息数据
	UserID := message.From.ID
	UserName := message.From.UserName
	text := message.Text
	// 是否发送消息触发器
	var sendMsg bool
	if message.Chat != nil {
		sendMsg = true
	}
	if message.Chat.IsGroup() && !strings.Contains(message.Text, "@"+bot.Self.UserName) {
		sendMsg = false //普通群组，被@才回复
	}
	if message.Chat.IsSuperGroup() && !strings.Contains(message.Text, "@"+bot.Self.UserName) {
		sendMsg = false //超级群组，被@才回复
	}

	if sendMsg {
		// 定义回复信息的数组
		replyMessages := []string{"你好,即将调用gpt3.5turbo的API"}

		if UserName == "viogami" {
			replyMessages[0] = "主人你好,即将为你调用gpt3.5turbo的API~"
		}

		// 调用ChatGPT API
		gptResponse, err := invokeChatGPTAPI(text)
		if err != nil {
			log.Printf("Error calling ChatGPT API: %v", err)
			gptResponse = "gpt调用失败了😥 错误信息：\n" + err.Error()
		}
		replyMessages = append(replyMessages, gptResponse)

		// 遍历发送每条信息
		for _, replymessage := range replyMessages {
			msg := tgbotapi.NewMessage(UserID, replymessage)
			msg.ReplyToMessageID = message.MessageID //@发信息的人回复
			_, err = bot.Send(msg)
			if err != nil {
				log.Println("Error sending message to user:", err)
			}
		}
	}
}
