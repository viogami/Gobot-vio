package tgbot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/viogami/Gobot-vio/config"
)

// 获取环境变量
var (
	BOT_TOKEN      = config.EnvConst.BOT_TOKEN
	TG_WEBHOOK_URL = config.EnvConst.TG_WEBHOOK_URL
)

var superUserId int64 //管理员id

func CreateTgbot() {
	// appConfig := flag.String("config", "./app.yaml", "application config path")
	// conf, _ := config.ConfigParse(*appConfig)
	// BOT_TOKEN = conf.Tgbot.BOT_TOKEN
	// TG_WEBHOOK_URL = conf.Tgbot.TG_WEBHOOK_URL

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
	// 此处的/tgbot/，是为了和URL匹配，URL：https://your-website/tgbot/
	updates := bot.ListenForWebhook("/tgbot/" + bot.Token)

	// 对监听到的updates遍历,并作出回应
	for update := range updates {
		log.Printf("get the message:%v", update.Message)
		if update.Message == nil {
			continue
		}
		//回复信息
		HandleIncomingMessage(bot, update.Message)
	}
}
