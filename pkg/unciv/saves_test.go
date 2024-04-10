package unciv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentPlayer(t *testing.T) {
	save, err := newSaveFromData(testSaveData)
	assert.Nil(t, err)
	currentPlayer, err := save.GetCurrentPlayer()
	assert.Nil(t, err)
	assert.Equal(t, Player{
		ChosenCiv:  "England",
		PlayerType: "Human",
		PlayerID:   "10000000-0000-0000-0000-000000000000",
	}, currentPlayer)
}
