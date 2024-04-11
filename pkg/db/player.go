package db

import (
	"fmt"
)

func (db *gormDB) GetPlayerByUncivID(uncivID string, getGames bool) (Player, error) {
	var player Player
	query := db.db.Model(&Player{})
	if getGames {
		query = query.Preload("Games")
	}
	err := query.First(&player, &Player{UncivID: uncivID}).Error
	if err != nil {
		return Player{}, fmt.Errorf("failed to get player: %w", err)
	}
	return player, nil
}
