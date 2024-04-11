package db

import (
	"fmt"
)

func (db *gormDB) GetGameByID(gameID string) (Game, error) {
	var game Game
	err := db.db.Model(&Game{}).Preload("Players").First(&game, &Game{GameID: gameID}).Error
	if err != nil {
		return Game{}, fmt.Errorf("failed to get game: %w", err)
	}

	return game, nil
}

func (db *gormDB) GetPlayerByUncivID(uncivID string) (Player, error) {
	var player Player
	err := db.db.Model(&Player{}).First(&player, &Player{UncivID: uncivID}).Error
	if err != nil {
		return Player{}, fmt.Errorf("failed to get player: %w", err)
	}
	return player, nil
}
