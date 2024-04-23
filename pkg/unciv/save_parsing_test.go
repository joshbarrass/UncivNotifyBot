package unciv

import (
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const gameParametersJSON = `{"difficulty":"Warlord","players":[{"chosenCiv":"Babylon","playerType":"Human","playerId":"00000000-0000-0000-0000-000000000000"},{"chosenCiv":"Rome","playerType":"Human","playerId":"00000000-0000-0000-0000-000000000000"},{"chosenCiv":"England","playerType":"Human","playerId":"10000000-0000-0000-0000-000000000000"}],"victoryTypes":["Scientific","Cultural","Domination","Diplomatic","Time"],"isOnlineMultiplayer":true,"multiplayerServerUrl":"https://uncivserver.xyz","baseRuleset":"Civ V - Vanilla"}`

var testCorrectURL, _ = url.Parse("https://uncivserver.xyz")
var testGameParameters = GameParameters{
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
			PlayerID:   "10000000-0000-0000-0000-000000000000",
		},
	},
	VictoryTypes:         []string{"Scientific", "Cultural", "Domination", "Diplomatic", "Time"},
	IsOnlineMultiplayer:  true,
	MultiplayerServerURL: testCorrectURL,
	BaseRuleset:          "Civ V - Vanilla",
}

const saveDataJSON = `{"version":{"number":3,"createdWith":{"text":"x.y.z","number":0}}, "gameParameters":` + gameParametersJSON + `,"turns":13,"currentPlayer":"England","currentTurnStartTime":1712768034170,"gameId":"00000000-0000-0000-0000-000000000000","historyStartTurn":0}`

var testSaveData = saveData{
	GameParameters:       testGameParameters,
	Turns:                13,
	CurrentFaction:       "England",
	CurrentTurnStartTime: time.Unix(1712768034170/1000, 0),
	GameID:               "00000000-0000-0000-0000-000000000000",
	HistoryStartTurn:     0,
}

func TestGameParametersUnmarshal(t *testing.T) {
	var s GameParameters
	err := json.Unmarshal([]byte(gameParametersJSON), &s)
	assert.Nil(t, err)
	assert.Equal(t, testGameParameters, s)
}

func TestSaveDataUnmarshal(t *testing.T) {
	var s saveData
	err := json.Unmarshal([]byte(saveDataJSON), &s)
	assert.Nil(t, err)
	assert.Equal(t, testSaveData, s)
}

func TestDownloadSave(t *testing.T) {
	var gameID = testSaveData.GameID
	serveMux := http.NewServeMux()
	serveMux.HandleFunc(SERVER_FILES_ROUTE+"/"+gameID, func(w http.ResponseWriter, _ *http.Request) {
		base64w := base64.NewEncoder(base64.StdEncoding, w)
		gzw := gzip.NewWriter(base64w)
		gzw.Write([]byte(saveDataJSON))
		gzw.Close()
		base64w.Close()
	})
	testServer := httptest.NewServer(serveMux)
	testClient := testServer.Client()

	server, err := NewUncivServer(testServer.URL, testClient)
	assert.Nil(t, err)

	s, err := server.DownloadSave(gameID)
	assert.Nil(t, err)
	expected, _ := newSaveFromData(testSaveData)
	assert.Equal(t, expected, s)
}
