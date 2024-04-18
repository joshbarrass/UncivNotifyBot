package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGameByID(t *testing.T) {
	DB := playerTestSetupDB(t)
	for i, expected := range []Game{TestGame1, TestGame2} {
		t.Run(fmt.Sprintf("Game%d", i+1), func(t *testing.T) {
			actual, err := DB.GetGameByID(expected.GameID, true)
			assert.Nil(t, err)
			// compare manually, since the times in the db make it trickier
			assert.Equal(t, expected.GameID, actual.GameID)
			assert.Equal(t, expected.ChatID, actual.ChatID)
			assert.Equal(t, len(expected.Players), len(actual.Players))
			for i, expectedPlayer := range expected.Players {
				assert.Equal(t, expectedPlayer.UncivID, actual.Players[i].UncivID)
				assert.Equal(t, expectedPlayer.TelegramID, actual.Players[i].TelegramID)
			}
		})
	}
}

func TestGetGameByChatID(t *testing.T) {
	DB := playerTestSetupDB(t)
	for i, expected := range []Game{TestGame1, TestGame2} {
		t.Run(fmt.Sprintf("Game%d", i+1), func(t *testing.T) {
			actual, err := DB.GetGameByChatID(expected.ChatID, true)
			assert.Nil(t, err)
			// compare manually, since the times in the db make it trickier
			assert.Equal(t, expected.GameID, actual.GameID)
			assert.Equal(t, expected.ChatID, actual.ChatID)
			assert.Equal(t, len(expected.Players), len(actual.Players))
			for i, expectedPlayer := range expected.Players {
				assert.Equal(t, expectedPlayer.UncivID, actual.Players[i].UncivID)
				assert.Equal(t, expectedPlayer.TelegramID, actual.Players[i].TelegramID)
			}
		})
	}
}

func TestAddGame(t *testing.T) {
	t.Run("Clean", func(t *testing.T) {
		DB, err := NewMemoryDB(true)
		assert.Nil(t, err)
		// var db = DB.(*gormDB)

		err = DB.AddGame(TestGame1)
		assert.Nil(t, err)
	})
	t.Run("Clash", func(t *testing.T) {
		DB := playerTestSetupDB(t)
		err := DB.AddGame(TestGame1)
		assert.NotEqual(t, nil, err)
	})
}
