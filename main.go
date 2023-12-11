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
	// 初始化bot
	bot, err := tgbotapi.NewBotAPI(BOT_TOKEN)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	//创建webhook,指向你的URL
	wh, _ := tgbotapi.NewWebhook(TG_WEBHOOK_URL + bot.Token)

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	//输出webhook信息,判断是否建立成功
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	// 监听webhook是否有更新,更新存放到updates中
	updates := bot.ListenForWebhook("/" + bot.Token)

	// 输出监听端口
	log.Println("Listenning on port", port, ".")
	go http.ListenAndServe(":"+port, nil)

	// 对监听到的updates)遍历,并作出回应
	for update := range updates {
		log.Printf("%+v\n", update)
		if update.Message == nil {
			continue
		}
		//回复信息
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(msg); err != nil {
			log.Fatal(err)
		}
	}
}
