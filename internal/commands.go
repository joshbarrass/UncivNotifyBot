package internal

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joshbarrass/UncivNotifyBot/pkg/db"
	"github.com/joshbarrass/UncivNotifyBot/pkg/telegrambot"
	"github.com/joshbarrass/UncivNotifyBot/pkg/unciv"
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

func CommandConnect(bot *telegrambot.Bot, update tgbotapi.Update) error {
	msg := telegrambot.GetMessageObject(update)
	userID := msg.From.ID
	args := telegrambot.CommandArgsSplit(update)
	if len(args) < 1 {
		bot.ReplyToMsg(update, MSG_ERR_BIND_NO_ARGS)
		return nil
	}
	uncivID := args[0]

	context := GetBotContext(bot)
	err := context.Database.ConnectTelegram(uncivID, int64(userID))
	if err != nil {
		logger := reportError(bot, update)
		logger.WithField(
			"args",
			args,
		).WithField(
			"userID",
			userID,
		).Errorf("failed to get human players: %s", err)
		return err
	}

	storedPlayer, err := context.Database.GetPlayerByUncivID(uncivID, false)
	if err != nil {
		// skip -- do nothing
		return nil
	}
	bot.ReplyToMsg(update, fmt.Sprintf("Found uncivID %s with telegramID %d", storedPlayer.UncivID, storedPlayer.TelegramID))
	return nil
}

func CommandBind(bot *telegrambot.Bot, update tgbotapi.Update) error {
	msg := telegrambot.GetMessageObject(update)
	chatID := msg.Chat.ID
	args := telegrambot.CommandArgsSplit(update)
	if len(args) < 1 {
		bot.ReplyToMsg(update, MSG_ERR_BIND_NO_ARGS)
		return nil
	}
	gameID := args[0]

	// test the gameID (and get the player list at the same time)
	uncivServer := unciv.NewDefaultUncivServer()
	save, err := uncivServer.DownloadSave(gameID)
	if err != nil {
		logger := reportError(bot, update)
		logger.WithField(
			"args",
			args,
		).WithField(
			"gameID",
			gameID,
		).Errorf("failed to download save: %s", err)
		return err
	}
	// extract the players
	players, err := save.GetHumanPlayers()
	if err != nil {
		logger := reportError(bot, update)
		logger.WithField(
			"args",
			args,
		).WithField(
			"gameID",
			gameID,
		).Errorf("failed to get human players: %s", err)
		return err
	}

	// convert the players from the save into objects we can store
	// in the DB
	dbPlayers := make([]db.Player, len(players))
	for i, player := range players {
		dbPlayers[i] = db.Player{
			UncivID: player.PlayerID,
		}
	}

	// create the new game
	game := db.Game{
		GameID:  gameID,
		ChatID:  chatID,
		Players: dbPlayers,
	}

	context := GetBotContext(bot)
	err = context.Database.AddGame(game)
	if err != nil {
		logger := reportError(bot, update)
		logger.WithField(
			"args",
			args,
		).WithField(
			"gameID",
			gameID,
		).Errorf("failed to get human players: %s", err)
		return err
	}

	storedGame, err := context.Database.GetGameByID(gameID, true)
	if err != nil {
		// skip -- do nothing
		return nil
	}
	bot.ReplyToMsg(update, fmt.Sprintf("Found game with ID %s (%d human players)", storedGame.GameID, len(storedGame.Players)))

	return nil
}
