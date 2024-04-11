package db

import (
	"time"

	"gorm.io/gorm"
)

type keylessModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Player struct {
	keylessModel
	UncivID    string `gorm:"primaryKey"`
	TelegramID int64
}

type Game struct {
	keylessModel
	GameID  string `gorm:"primaryKey"`
	ChatID  int64
	Players []Player `gorm:"many2many:game_players;"`
}
