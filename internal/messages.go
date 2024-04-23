package internal

const (
	MSG_START_FMT          = "Hi %s! This is the Unciv Notification Bot!"
	MSG_ERR_NOT_FOUND      = "Sorry, I couldn't find that command. Please try something else."
	MSG_ERR_UNEXPECTED_FMT = "An unexpected error has occurred. Please try again later.\n\nError reference: %s" // error reference
	MSG_ERR_NO_GAME_FOUND  = "Did not find any games bound to this chat. Use the /bind command with your game's ID to bind a game to this chat."
)

const (
	MSG_TURN_FMT              = "It is your turn! (as of %s)"
	MSG_TURN_REGISTERED_FMT   = "[%s](tg://user?id=%d): " + MSG_TURN_FMT
	MSG_TURN_UNREGISTERED_FMT = "%s: " + MSG_TURN_FMT
)

const (
	MSG_ERR_BIND_NO_ARGS    = "Syntax: /bind <game ID>"
	MSG_ERR_CONNECT_NO_ARGS = "Syntax: /register <unciv ID>"
)
