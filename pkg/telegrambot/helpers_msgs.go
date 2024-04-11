package telegrambot

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// GetMessageObject will return the main message object from an
// update. This excludes the message objects for when a message has
// been edited. If there are no main message objects (for example, if
// the update is an InlineQuery), then nil is returned.
func GetMessageObject(update tgbotapi.Update) *tgbotapi.Message {
	if update.Message != nil {
		return update.Message
	}
	if update.ChannelPost != nil {
		return update.ChannelPost
	}
	return nil
}

// GetUpdateMessageID extracts the ID of the message that triggered an
// update. If the update does not contain a message ID, ErrNoMessageID
// is returned.
func GetUpdateMessageID(update tgbotapi.Update) (int, error) {
	if msg := GetMessageObject(update); msg != nil {
		return msg.MessageID, nil
	}
	return 0, ErrNoMessageID
}

// ReplyToMsg sends a message of a given text as a reply to the
// message in the update. The Markdown parser is assumed. If the
// update does not contain a message that can be replied to, the
// message is sent without a reply.
func (bot *Bot) ReplyToMsg(update tgbotapi.Update, text string) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	replyToID, err := GetUpdateMessageID(update)
	if err == nil {
		msg.ReplyToMessageID = replyToID
	}

	msg.ParseMode = tgbotapi.ModeMarkdown
	bot.API.Send(msg)
	return nil
}

// CommandArgs returns the arguments to a command as a single
// string. Alias of update.ChannelPost.CommandArguments
func CommandArgs(update tgbotapi.Update) string {
	msg := GetMessageObject(update)
	if msg == nil {
		return ""
	}
	return msg.CommandArguments()
}

// CommandArgsSplit returns the arguments to a command, split by space
func CommandArgsSplit(update tgbotapi.Update) []string {
	msg := GetMessageObject(update)
	if msg == nil {
		return []string{}
	}
	return strings.Split(msg.CommandArguments(), " ")
}
