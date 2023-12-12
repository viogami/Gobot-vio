package main

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
		//定义回复字段
		ResponseText := ""

		if userID == 5094809802 {
			ResponseText = "主人你好~\n"
		}

		// 调用ChatGPT API
		gptResponse, err := invokeChatGPTAPI(text)
		if err != nil {
			log.Printf("Error calling ChatGPT API: %v", err)
			return
		}

		// 发送ChatGPT的回复给用户
		msg := tgbotapi.NewMessage(userID, ResponseText+gptResponse)
		msg.ReplyToMessageID = message.MessageID //@发信息的人回复
		_, err = bot.Send(msg)
		if err != nil {
			log.Println("Error sending message to user:", err)
		}
	}
}
