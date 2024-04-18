package internal

const (
	MSG_START_FMT          = "Hi %s! This is the Unciv Notification Bot!"
	MSG_ERR_NOT_FOUND      = "Sorry, I couldn't find that command. Please try something else."
	MSG_ERR_UNEXPECTED_FMT = "An unexpected error has occurred. Please try again later.\n\nError reference: %s" // error reference
)

const (
	MSG_TURN                  = "It is your turn!"
	MSG_TURN_REGISTERED_FMT   = "[%s](tg://user?id=%d): " + MSG_TURN
	MSG_TURN_UNREGISTERED_FMT = "%s: " + MSG_TURN
)

const (
	MSG_ERR_BIND_NO_ARGS    = "Syntax: /bind <game ID>"
	MSG_ERR_CONNECT_NO_ARGS = "Syntax: /register <unciv ID>"
)
