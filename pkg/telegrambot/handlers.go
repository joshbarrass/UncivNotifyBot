package telegrambot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type HandlerFunc func(*Bot, tgbotapi.Update) error

// Handler represents a way of handling an update. It must implement:
//   Filter() *Filter
//     This method returns the filter responsible for checking
//     whether the handler applies to this update. If the filter
//     returns true, Handle will be called with the update as its arg
//   Handle(*Bot, tgbotapi.Update) error
//     This method is responsible for handling the update, including
//     sending of any messages.
type Handler interface {
	Filter() Filter
	Handle(*Bot, tgbotapi.Update) error
}

// Filter represents a generic, reusable object for checking whether
// an update applies to a handler.
type Filter interface {
	Test(tgbotapi.Update) bool
}

// checks whether an update is:
//  1) a message
//  2) a matching command
type commandFilter struct {
	Command string
}

func newCommandFilter(command string) *commandFilter {
	return &commandFilter{
		Command: command,
	}
}

func (f *commandFilter) Test(update tgbotapi.Update) bool {
	msg := GetMessageObject(update)
	if msg == nil {
		return false
	}
	return msg.Command() == f.Command
}

// CommandHandler runs a particular function in response to a
// particular command, prefixed with a slash.
type CommandHandler struct {
	handler HandlerFunc
	filter  *commandFilter
}

func (handler *CommandHandler) Filter() Filter {
	return handler.filter
}

func (handler *CommandHandler) Handle(bot *Bot, update tgbotapi.Update) error {
	return handler.handler(bot, update)
}

// NewCommandHandler creates a new handler for a particular command.
func NewCommandHandler(command string, handler HandlerFunc) *CommandHandler {
	return &CommandHandler{
		handler: handler,
		filter:  newCommandFilter(command),
	}
}

// AllMessageHandler is a handler that matches all message-like updates.
type AllMessageHandler struct {
	handler HandlerFunc
	filter  *allMessageFilter
}

// allMessageFilter is a filter that matches all message-like updates.
type allMessageFilter struct{}

func (_ *allMessageFilter) Test(update tgbotapi.Update) bool {
	msg := GetMessageObject(update)
	return msg != nil
}

func (handler *AllMessageHandler) Filter() Filter {
	return handler.filter
}

func (handler *AllMessageHandler) Handle(bot *Bot, update tgbotapi.Update) error {
	return handler.handler(bot, update)
}

// NewAllMessageHandler creates a new handler that matches all message-like updates.
func NewAllMessageHandler(handler HandlerFunc) *AllMessageHandler {
	return &AllMessageHandler{
		handler: handler,
		filter:  &allMessageFilter{},
	}
}
