package message

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleIncomingMessage 处理用户消息
func HandleIncomingMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	// 分析消息数据
	userID := message.From.ID
	text := message.Text
	// 是否发送消息触发器
	sendMsg := true
	if message.Chat == nil {
		sendMsg = false
	}
	if message.Chat.IsGroup() && !strings.Contains(message.Text, "@"+bot.Self.UserName) {
		sendMsg = false
	}

	if sendMsg {
		// 定义回复信息的数组
		replyMessages := []string{"你好,即将调用gpt3.5turbo的API"}

		if userID == 5094809802 {
			replyMessages[0] = "主人你好,即将为你调用gpt3.5turbo的API~"
		}

		// 调用ChatGPT API
		gptResponse, err := invokeChatGPTAPI(text)
		if err != nil {
			log.Printf("Error calling ChatGPT API: %v", err)
			return
		}
		replyMessages = append(replyMessages, gptResponse)

		// 遍历发送每条信息
		for _, replymessage := range replyMessages {
			msg := tgbotapi.NewMessage(userID, replymessage)
			msg.ReplyToMessageID = message.MessageID //@发信息的人回复
			_, err = bot.Send(msg)
			if err != nil {
				log.Println("Error sending message to user:", err)
			}
		}
	}
}
