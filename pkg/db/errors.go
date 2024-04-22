package db

import "errors"

var (
	ErrUncivIDAlreadyBound = errors.New("Unciv ID already connected to a Telegram ID")
	ErrGameNotFound        = errors.New("game not found in database")
)
