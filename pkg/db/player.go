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

func (db *gormDB) AddPlayer(player Player) error {
	return db.db.Create(&player).Error
}

// ConnectTelegram binds a telegram user ID to an Unciv ID. If the
// Unciv ID is not already in the database, it will be created. If the
// Unciv ID is already in the database, the binding will be added if
// there is not already a telegram ID bound to it.
func (db *gormDB) ConnectTelegram(uncivID string, telegramID int64) error {
	var player Player
	if err := db.db.Model(&Player{}).Where(&Player{UncivID: uncivID}).FirstOrCreate(&player).Error; err != nil {
		return fmt.Errorf("unable to get player from db: %w", err)
	}
	if player.TelegramID != 0 {
		return ErrUncivIDAlreadyBound
	}
	if err := db.db.Model(&player).Update("TelegramID", telegramID).Error; err != nil {
		return fmt.Errorf("unable to update telegram ID: %w", err)
	}
	return nil
}
