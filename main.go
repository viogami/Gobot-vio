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
	chatGPTAPIURL  = "https://api.openai.com/v1/completions"
	chatGPTAPIKey  = os.Getenv("chatGPTAPIKey")
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
	go http.ListenAndServe(":8443", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
