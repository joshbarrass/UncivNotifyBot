package unciv

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const gameParametersJSON = `{"difficulty":"Warlord","players":[{"chosenCiv":"Babylon","playerType":"Human","playerId":"00000000-0000-0000-0000-000000000000"},{"chosenCiv":"Rome","playerType":"Human","playerId":"00000000-0000-0000-0000-000000000000"},{"chosenCiv":"England","playerType":"Human","playerId":"00000000-0000-0000-0000-000000000000"}],"victoryTypes":["Scientific","Cultural","Domination","Diplomatic","Time"],"isOnlineMultiplayer":true,"multiplayerServerUrl":"https://uncivserver.xyz","baseRuleset":"Civ V - Vanilla"}`

const saveDataJSON = `{"gameParameters":` + gameParametersJSON + `,"turns":13,"currentPlayer":"England","currentTurnStartTime":1712768034170,"gameId":"00000000-0000-0000-0000-000000000000","historyStartTurn":0}`

func TestGameParametersUnmarshal(t *testing.T) {
	var s GameParameters
	err := json.Unmarshal([]byte(gameParametersJSON), &s)
	assert.Nil(t, err)
	correctURL, _ := url.Parse("https://uncivserver.xyz")
	assert.Equal(t, GameParameters{
		Difficulty: "Warlord",
		Players: []Player{
			{
				ChosenCiv:  "Babylon",
				PlayerType: "Human",
				PlayerID:   "00000000-0000-0000-0000-000000000000",
			},
			{
				ChosenCiv:  "Rome",
				PlayerType: "Human",
				PlayerID:   "00000000-0000-0000-0000-000000000000",
			},
			{
				ChosenCiv:  "England",
				PlayerType: "Human",
				PlayerID:   "00000000-0000-0000-0000-000000000000",
			},
		},
		VictoryTypes:         []string{"Scientific", "Cultural", "Domination", "Diplomatic", "Time"},
		IsOnlineMultiplayer:  true,
		MultiplayerServerURL: correctURL,
		BaseRuleset:          "Civ V - Vanilla",
	}, s)
}

func TestSaveDataUnmarshal(t *testing.T) {
	var s saveData
	err := json.Unmarshal([]byte(saveDataJSON), &s)
	assert.Nil(t, err)
	correctURL, _ := url.Parse("https://uncivserver.xyz")
	assert.Equal(t, saveData{
		GameParameters: GameParameters{
			Difficulty: "Warlord",
			Players: []Player{
				{
					ChosenCiv:  "Babylon",
					PlayerType: "Human",
					PlayerID:   "00000000-0000-0000-0000-000000000000",
				},
				{
					ChosenCiv:  "Rome",
					PlayerType: "Human",
					PlayerID:   "00000000-0000-0000-0000-000000000000",
				},
				{
					ChosenCiv:  "England",
					PlayerType: "Human",
					PlayerID:   "00000000-0000-0000-0000-000000000000",
				},
			},
			VictoryTypes:         []string{"Scientific", "Cultural", "Domination", "Diplomatic", "Time"},
			IsOnlineMultiplayer:  true,
			MultiplayerServerURL: correctURL,
			BaseRuleset:          "Civ V - Vanilla",
		},
		Turns:                13,
		CurrentPlayer:        "England",
		CurrentTurnStartTime: time.Unix(1712768034170, 0),
		GameID:               "00000000-0000-0000-0000-000000000000",
		HistoryStartTurn:     0,
	}, s)
}
