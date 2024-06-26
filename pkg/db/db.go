package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const memoryDatabaseFilename = ":memory:?cache=shared"

type Database interface {
	GetGameByID(string, bool) (Game, error)
	GetGameByChatID(int64, bool) (Game, error)
	GetPlayerByUncivID(string, bool) (Player, error)
	AddGame(Game) error
	AddPlayer(Player) error
	ConnectTelegram(string, int64) error
}

type gormDB struct {
	db *gorm.DB
}

func NewSqliteDB(filename string, automigrate bool) (Database, error) {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if automigrate {
		db.AutoMigrate(&Player{}, &Game{})
	}
	return &gormDB{
		db: db,
	}, nil
}

func NewMemoryDB(automigrate bool) (Database, error) {
	return NewSqliteDB(memoryDatabaseFilename, automigrate)
}
