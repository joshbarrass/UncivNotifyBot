package db

import (
	"fmt"
)

func (db *gormDB) GetPlayerByUncivID(uncivID string) (Player, error) {
	var player Player
	err := db.db.Model(&Player{}).First(&player, &Player{UncivID: uncivID}).Error
	if err != nil {
		return Player{}, fmt.Errorf("failed to get player: %w", err)
	}
	return player, nil
}
