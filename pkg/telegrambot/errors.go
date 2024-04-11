package telegrambot

import "errors"

var (
	ErrNoMessageID = errors.New("update does not contain a message ID")
	ErrNoUsername  = errors.New("user does not have a username")
)
