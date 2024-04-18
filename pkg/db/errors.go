package db

import "errors"

var (
	ErrUncivIDAlreadyBound = errors.New("Unciv ID already connected to a Telegram ID")
)
