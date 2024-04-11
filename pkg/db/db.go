package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const memoryDatabaseFilename = ":memory:?cache=shared"

type Database interface {
	GetPlayersByGameID(string) ([]Player, error)
	GetPlayerByUncivID(string) (Player, error)
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
