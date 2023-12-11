package main

import (
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// 定义环境变量
var (
	BOT_TOKEN      = os.Getenv("BOT_TOKEN")
	TG_WEBHOOK_URL = os.Getenv("TG_WEBHOOK_URL")
	chatGPTAPIKey  = os.Getenv("chatGPTAPIKey")
	port           = os.Getenv("PORT")
)

const (
	chatGPTAPIURL = "https://api.openai.com/v1/completions"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(BOT_TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	wh, _ := tgbotapi.NewWebhook(TG_WEBHOOK_URL + bot.Token)

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	log.Println("Listenning on port", port, ".")
	go http.ListenAndServe(":"+port, nil)

	for update := range updates {
		log.Printf("%+v\n", update)
		if update.Message == nil {
			// ignore any non-Message Updates
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(msg); err != nil {
			log.Fatal(err)
		}
	}
}
