package main

import (
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// 读取环境变量
	BOT_TOKEN := os.Getenv("BOT_TOKEN")
	webhookURL := os.Getenv("TELEGRAM_WEBHOOK_URL")
	if webhookURL == "" || BOT_TOKEN == "" {
		log.Println("env var not set!")
	}

	// 创建一个 Telegram Bot 实例
	bot, err := tgbotapi.NewBotAPI(BOT_TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	// 启用调试模式
	bot.Debug = true

	// 打印 Bot 用户名，表示授权成功
	log.Printf("成功授权给： %s", bot.Self.UserName)

	// 创建一个 Webhook
	wh, _ := tgbotapi.NewWebhook(webhookURL)

	// 使用 Bot 实例向 Telegram 设置 Webhook
	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	// 获取 Webhook 信息，检查是否设置成功
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	// 如果 Webhook 设置失败，打印错误信息
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback 失败: %s", info.LastErrorMessage)
	}

	// 监听来自 Telegram 的 Webhook 更新
	updates := bot.ListenForWebhook("/" + bot.Token)
	log.Println(updates)

	// 启动 HTTPS 服务器，用于接收 Telegram 的 Webhook 更新
	go http.ListenAndServeTLS("0.0.0.0:443", "cert.pem", "key.pem", nil)
	log.Println("启动 HTTPS 服务器，用于接收 Telegram 的 Webhook 更新")

	// 循环处理来自 Telegram 的更新
	for update := range updates {
		// 检查更新是否为消息类型
		if update.Message == nil {
			continue
		}

		// 打印收到的消息文本
		log.Printf("收到消息: %s", update.Message.Text)

		// 构建回复消息
		reply := tgbotapi.NewMessage(update.Message.Chat.ID, "我收到你的消息啦~")

		// 发送回复消息
		_, err := bot.Send(reply)
		if err != nil {
			log.Println("发送回复消息失败:", err)
		}
	}
}
