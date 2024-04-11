package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestPlayer1 = Player{
	UncivID:    "00000000-0000-0000-0000-000000000000",
	TelegramID: 1,
}
var TestPlayer2 = Player{
	UncivID:    "10000000-0000-0000-0000-000000000000",
	TelegramID: 2,
}
var TestPlayer3 = Player{
	UncivID:    "20000000-0000-0000-0000-000000000000",
	TelegramID: 2,
}
var TestGame1 = Game{
	GameID:  "00000000-0000-0000-0000-000000000000",
	ChatID:  42,
	Players: []Player{TestPlayer1, TestPlayer2},
}
var TestGame2 = Game{
	GameID:  "10000000-0000-0000-0000-000000000000",
	Players: []Player{TestPlayer3},
}

func playerTestSetupDB(t *testing.T) *gormDB {
	DB, err := NewSqliteDB(memoryDatabaseFilename, true)
	var db = DB.(*gormDB)
	assert.Nil(t, err)
	assert.Nil(t, db.db.Create(&TestPlayer1).Error)
	assert.Nil(t, db.db.Create(&TestPlayer2).Error)
	assert.Nil(t, db.db.Create(&TestPlayer3).Error)
	assert.Nil(t, db.db.Create(&TestGame1).Error)
	assert.Nil(t, db.db.Create(&TestGame2).Error)
	return db
}

func TestGetPlayerByUncivID(t *testing.T) {
	DB := playerTestSetupDB(t)
	for i, expected := range []Player{TestPlayer1, TestPlayer2, TestPlayer3} {
		t.Run(fmt.Sprintf("Player%d", i+1), func(t *testing.T) {
			actual, err := DB.GetPlayerByUncivID(expected.UncivID)
			assert.Nil(t, err)
			// compare manually, since the times in the db make it trickier
			assert.Equal(t, expected.UncivID, actual.UncivID)
			assert.Equal(t, expected.TelegramID, actual.TelegramID)
		})
	}
}
