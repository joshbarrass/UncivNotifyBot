package telegrambot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// GetFromUser returns the User from whom the update was received,
// regardless of the true underlying type of the update.
func GetFromUser(update tgbotapi.Update) *tgbotapi.User {
	var u *tgbotapi.User

	switch {
	case update.Message != nil:
		u = update.Message.From
	case update.EditedMessage != nil:
		u = update.EditedMessage.From
	case update.ChannelPost != nil:
		u = update.ChannelPost.From
	case update.EditedChannelPost != nil:
		u = update.EditedChannelPost.From
	case update.InlineQuery != nil:
		u = update.InlineQuery.From
	case update.ChosenInlineResult != nil:
		u = update.ChosenInlineResult.From
	case update.CallbackQuery != nil:
		u = update.CallbackQuery.From
	case update.ShippingQuery != nil:
		u = update.ShippingQuery.From
	case update.PreCheckoutQuery != nil:
		u = update.PreCheckoutQuery.From
	}

	return u
}

// GetUserFirstName returns the first name of the user who sent this
// update. There is no error returned for this since a first name is
// required in Telegram.
func GetUserFirstName(update tgbotapi.Update) string {
	return GetFromUser(update).FirstName
}

// GetUserID returns the ID of the user who sent this update. There is
// no error returned for this since a first name is required.
func GetUserID(update tgbotapi.Update) int {
	return GetFromUser(update).ID
}

// GetUsername returns the ID of the user who sent this update. If the
// user does not have a username, ErrNoUsername is returned.
func GetUsername(update tgbotapi.Update) (string, error) {
	user := GetFromUser(update)
	if user.UserName == "" {
		return "", ErrNoUsername
	}
	return user.UserName, nil
}
