package internal

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joshbarrass/UncivNotifyBot/pkg/telegrambot"
)

const (
	MSG_START_FMT     = "Hi %s! This is the Unciv Notification Bot!"
	MSG_ERR_NOT_FOUND = "Sorry, I couldn't find that command. Please try something else."
)

func CommandStart(bot *telegrambot.Bot, update tgbotapi.Update) error {
	userForename := telegrambot.GetUserFirstName(update)
	bot.ReplyToMsg(update, fmt.Sprintf(MSG_START_FMT, userForename))
	return nil
}

func CommandNotFound(bot *telegrambot.Bot, update tgbotapi.Update) error {
	// TODO: maybe determine whether we are in a channel/group
	// chat where the bot could have been messaged by mistake, and
	// ignore if that is the case.
	bot.ReplyToMsg(update, MSG_ERR_NOT_FOUND)
	return nil
}
