package main

import (
	"github.com/joshbarrass/UncivNotifyBot/internal"
	"github.com/joshbarrass/UncivNotifyBot/pkg/telegrambot"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Configuration struct {
	BotToken  string `envconfig:"TOKEN" required:"true"`
	DebugLogs bool   `envconfig:"DEBUG_LOGS" default:"0"`
	PoolSize  int    `envconfig:"POOL_SIZE" default:"4"`
	DBPath    string `envconfig:"DB_PATH" default:""`
}

func main() {
	var config Configuration
	err := envconfig.Process("UNCIVBOT", &config)
	if err != nil {
		logrus.Fatalf("Failed to process config: %s", err)
	}

	if config.DebugLogs {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("debug logs enabled")
	}

	bot, err := telegrambot.NewBot(config.BotToken)
	if err != nil {
		logrus.Fatalf("Failed to initialise bot: %s", err)
	}

	logrus.Debug("Creating context...")
	context := internal.NewBotContext(bot)
	if config.DBPath == "" {
		logrus.Debug("Initialising memory DB...")
		err = context.InitialiseMemoryDB()
		if err != nil {
			logrus.Fatalf("Failed to initialise memory DB: %s", err)
		}
	} else {
		logrus.Debug("Initialising sqlite DB...")
		err = context.InitialiseSQLiteDB(config.DBPath)
		if err != nil {
			logrus.Fatalf("Failed to initialise sqlite DB: %s", err)
		}
	}
	logrus.Debug("Initialised DB")
	logrus.Debug("Created context!")

	logrus.Debug("Adding handlers...")
	bot.AddHandler(telegrambot.NewCommandHandler("start", internal.CommandStart))
	bot.AddHandler(telegrambot.NewCommandHandler("bind", internal.CommandBind))
	bot.AddHandler(telegrambot.NewCommandHandler("register", internal.CommandConnect))
	bot.AddHandler(telegrambot.NewCommandHandler("turn", internal.CommandTurn))

	// catch-all error message
	bot.AddHandler(telegrambot.NewAllMessageHandler(internal.CommandNotFound))

	logrus.Debug("Finished adding handlers!")

	logrus.Debug("Starting update handler...")
	bot.HandleUpdates(config.PoolSize)
}
