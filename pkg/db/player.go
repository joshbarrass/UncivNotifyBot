package db

import (
	"fmt"
)

func (db *gormDB) GetPlayersByGameID(gameID string) ([]Player, error) {
	// get the user IDs for the game from the DB
	var game Game
	err := db.db.Model(&Game{}).Preload("Players").First(&game, &Game{GameID: gameID}).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get game: %w", err)
	}

	return game.Players, nil
}
