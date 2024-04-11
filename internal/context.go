package internal

import (
	"github.com/joshbarrass/UncivNotifyBot/pkg/db"
	"github.com/joshbarrass/UncivNotifyBot/pkg/telegrambot"
)

// BotContext is used to store additional information in the bot object
type BotContext struct {
	Database db.Database
}

func NewBotContext(bot *telegrambot.Bot) *BotContext {
	context := &BotContext{}
	bot.SetContext(context)
	return context
}

func GetBotContext(bot *telegrambot.Bot) *BotContext {
	return bot.GetContext().(*BotContext)
}

func (context *BotContext) InitialiseMemoryDB() (err error) {
	context.Database, err = db.NewMemoryDB(true)
	return
}
