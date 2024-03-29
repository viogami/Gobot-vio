package tgbot

import (
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/**
 * 检查是否是群组的管理员
 */
func checkAdmin(bot *tgbotapi.BotAPI, gid int64, user tgbotapi.User) bool {
	admins, _ := bot.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID: gid}})
	uid := user.ID
	if superUserId > 0 && uid == superUserId {
		return true
	}
	for _, user := range admins {
		if uid == user.User.ID {
			return true
		}
	}
	return false
}

/*
 * 禁言群员
 */
func banMember(bot *tgbotapi.BotAPI, gid int64, uid int64, sec int64) {
	if sec <= 0 {
		sec = 9999999999999
	}
	banChatMemberConfig := tgbotapi.BanChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID: gid,
			UserID: uid,
		},
		UntilDate:      time.Now().Unix() + sec,
		RevokeMessages: false,
	}
	_, _ = bot.Request(banChatMemberConfig)
}

func unbanMember(bot *tgbotapi.BotAPI, gid int64, uid int64) {
	banChatMemberConfig := tgbotapi.BanChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID: gid,
			UserID: uid,
		},
		UntilDate:      9999999999999,
		RevokeMessages: false,
	}
	_, _ = bot.Request(banChatMemberConfig)
}

/**
 * 踢出群员
 */
func kickMember(bot *tgbotapi.BotAPI, gid int64, uid int64) {
	kickChatMemberConfig := tgbotapi.KickChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID: gid,
			UserID: uid,
		},
	}
	_, _ = bot.Request(kickChatMemberConfig)
}

/**
 * 返回群组的所有管理员, 用来进行一次性@
 */
func getAdmins(bot *tgbotapi.BotAPI, gid int64) string {
	admins, _ := bot.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID: gid}})
	list := ""
	for _, admin := range admins {
		user := admin.User
		if user.IsBot {
			continue
		}
		list += "[" + user.String() + "](tg://user?id=" + strconv.FormatInt(admin.User.ID, 10) + ")\r\n"
	}
	return list
}
