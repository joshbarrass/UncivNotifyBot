package internal

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
	"github.com/joshbarrass/UncivNotifyBot/pkg/telegrambot"
	"github.com/sirupsen/logrus"
)

// generate a unique error reference based on a UUID
func generateErrorReference() string {
	u := uuid.New()
	raw := strings.ReplaceAll(u.String(), "-", "")
	data, _ := hex.DecodeString(raw)
	return base64.StdEncoding.EncodeToString(data)
}

// sends a message to the user indicating that an error occurred and
// providing a reference. A logger is then returned for further
// logging with this reference.
func reportError(bot *telegrambot.Bot, update tgbotapi.Update) *logrus.Entry {
	ref := generateErrorReference()
	msg := fmt.Sprintf(MSG_ERR_UNEXPECTED_FMT, ref)
	bot.ReplyToMsg(update, msg, true)
	return logrus.WithField("reference", ref)
}
