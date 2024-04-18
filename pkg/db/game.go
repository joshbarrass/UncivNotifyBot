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

func (db *gormDB) AddGame(game Game) error {
	return db.db.Create(&game).Error
}

func (db *gormDB) GetGameByChatID(chatID int64, getPlayers bool) (Game, error) {
	var game Game
	query := db.db.Model(&Game{})
	if getPlayers {
		query.Preload("Players")
	}
	err := query.First(&game, &Game{ChatID: chatID}).Error
	if err != nil {
		return Game{}, fmt.Errorf("failed to get game: %w", err)
	}

	return game, nil
}
