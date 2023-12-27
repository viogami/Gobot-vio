package tgbot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/*
 * 发送文字消息
 */

/*
 * 发送图片消息, 需要是已经存在的图片链接
 */
func SendPhoto(chatid int64, photoid tgbotapi.RequestFileData) tgbotapi.Message {

	file := tgbotapi.NewPhoto(chatid, photoid)
	mmsg, err := bot.Send(file)
	if err != nil {
		log.Println(err)
	}

	return mmsg
}

/*
 * 发送动图, 需要是已经存在的链接
 */
func sendGif(chatid int64, gifid tgbotapi.RequestFileData) tgbotapi.Message {
	file := tgbotapi.NewAnimation(chatid, gifid)
	mmsg, err := bot.Send(file)
	if err != nil {
		log.Println(err)
	}

	return mmsg
}

/*
 * 发送视频, 需要是已经存在的视频连接
 */
func sendVideo(chatid int64, videoid tgbotapi.RequestFileData) tgbotapi.Message {
	file := tgbotapi.NewVideo(chatid, videoid)
	mmsg, err := bot.Send(file)
	if err != nil {
		log.Println(err)
	}

	return mmsg
}

/*
 * 发送文件, 需要是已经存在的文件链接
 */
func sendFile(chatid int64, fileid tgbotapi.RequestFileData) tgbotapi.Message {
	file := tgbotapi.NewDocument(chatid, fileid)
	mmsg, err := bot.Send(file)
	if err != nil {
		log.Println(err)
	}

	return mmsg
}
