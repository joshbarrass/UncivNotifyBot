package telegrambot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const MaxUint uint = ^uint(0)
const MaxInt int = int(MaxUint >> 1)

type empty struct{}

// Bot represents the bot itself, stores the necessary variables for
// the functioning of the bot, and exposes the necessary methods for
// the bot's operation.
type Bot struct {
	API *tgbotapi.BotAPI
	//
	handlers []Handler
	//
	handlerPool chan empty
	done        chan empty

	// context stores an object that holds application-specific
	// information. It can be any user-defined object, and then
	// must be cast back to the original type
	context interface{}
}

func NewBot(apiKey string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return nil, err
	}
	return &Bot{
		API:  api,
		done: make(chan empty, 1),
	}, nil
}

// AddHandler adds a new update handler to the bot. The order in which
// handlers are added sets their priorities; handlers are checked to
// see whether an update applies in the order they are added here, so
// if a filter could match multiple commands, this order is important.
func (bot *Bot) AddHandler(handler Handler) {
	bot.handlers = append(bot.handlers, handler)
}

// GetUpdatesChan returns the updates channel for the bot, using the bot's stored API.
func (bot *Bot) GetUpdatesChan() (tgbotapi.UpdatesChannel, error) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	return bot.API.GetUpdatesChan(updateConfig)
}

// HandleUpdates is the main method for the bot. It begins the loop of
// handling updates.
func (bot *Bot) HandleUpdates(poolSize int) error {
	// create and fill the handler pool
	bot.handlerPool = make(chan empty, poolSize)
	for i := 0; i < poolSize; i += 1 {
		bot.handlerPool <- empty{}
	}

	// create the update channel
	updates, err := bot.GetUpdatesChan()
	if err != nil {
		return fmt.Errorf("failed to create update channel: %w", err)
	}

	// loop indefinitely
	for {
		select {
		case update := <-updates:
			// wait for a handler to be available
			<-bot.handlerPool
			// launch a handler as a goroutine
			go func() {
				bot.handleUpdate(update)
				bot.handlerPool <- empty{}
			}()
		case <-bot.done:
			// wait for all handlers to finish and return to the pool
			for i := 0; i < poolSize; i += 1 {
				<-bot.handlerPool
			}
			bot.API.StopReceivingUpdates()
			return nil
		}
	}
}

// Shutdown waits for all handlers to finish and then gracefully shuts
// down the update handler and exits.
func (bot *Bot) Shutdown() {
	bot.done <- empty{}
}

func (bot *Bot) handleUpdate(update tgbotapi.Update) {
	// loop through the handlers checking whether they apply, and
	// call the first that does
	for _, handler := range bot.handlers {
		if handler.Filter().Test(update) {
			handler.Handle(bot, update)
			break
		}
	}
}

// SetContext stores an object that holds application-specific
// information in the bot object. It can be any user-defined object,
// and then must be cast back to the original type.
func (bot *Bot) SetContext(context interface{}) {
	bot.context = context
}

// GetContext returns the stored context object, ready for the user to
// cast back to its original type.
func (bot *Bot) GetContext() interface{} {
	return bot.context
}
