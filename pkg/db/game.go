package db

import "fmt"

func (db *gormDB) GetGameByID(gameID string, getPlayers bool) (Game, error) {
	var game Game
	query := db.db.Model(&Game{})
	if getPlayers {
		query.Preload("Players")
	}
	err := query.First(&game, &Game{GameID: gameID}).Error
	if err != nil {
		return Game{}, fmt.Errorf("failed to get game: %w", err)
	}

	return game, nil
}
