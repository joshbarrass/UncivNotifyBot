package db

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (db *gormDB) GetGameByID(gameID string, getPlayers bool) (Game, error) {
	var game Game
	query := db.db.Model(&Game{})
	if getPlayers {
		query.Preload("Players")
	}
	err := query.First(&game, &Game{GameID: gameID}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Game{}, fmt.Errorf("failed to get game by ID: %w (%w)", ErrGameNotFound, err)
	}
	if err != nil {
		return Game{}, fmt.Errorf("failed to get game by ID: %w", err)
	}

	return game, nil
}

func (db *gormDB) AddGame(game Game) error {
	return db.db.Create(&game).Error
}

// TODO: should be GetGamesByChatID
func (db *gormDB) GetGameByChatID(chatID int64, getPlayers bool) (Game, error) {
	var game Game
	query := db.db.Model(&Game{})
	if getPlayers {
		query.Preload("Players")
	}
	err := query.First(&game, &Game{ChatID: chatID}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Game{}, fmt.Errorf("failed to get game by chat ID: %w (%w)", ErrGameNotFound, err)
	}
	if err != nil {
		return Game{}, fmt.Errorf("failed to get game by chat ID: %w", err)
	}

	return game, nil
}
