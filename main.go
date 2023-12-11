package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	BOT_TOKEN := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(BOT_TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			text := "你发什么我都喜欢你"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
