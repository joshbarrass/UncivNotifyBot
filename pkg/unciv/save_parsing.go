package unciv

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type saveData struct {
	GameParameters       GameParameters
	Turns                int
	CurrentFaction       string
	CurrentTurnStartTime time.Time
	GameID               string
	HistoryStartTurn     int
}

func (sd *saveData) UnmarshalJSON(data []byte) error {
	// unmarshal to intermediate first
	var intermediate saveDataIntermediate
	err := json.Unmarshal(data, &intermediate)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to intermediate: %w", err)
	}

	err = intermediate.Convert(sd)
	if err != nil {
		return fmt.Errorf("failed to convert intermediate: %w", err)
	}
	return nil
}

type saveDataIntermediate struct {
	GameParameters       gameParametersIntermediate `json:"gameParameters"`
	Turns                int                        `json:"turns"`
	CurrentPlayer        string                     `json:"currentPlayer"`
	CurrentTurnStartTime int64                      `json:"currentTurnStartTime"`
	GameID               string                     `json:"gameId"`
	HistoryStartTurn     int                        `json:"historyStartTurn"`
}

func (data saveDataIntermediate) Convert(v *saveData) error {
	var gameParameters GameParameters
	err := data.GameParameters.Convert(&gameParameters)
	if err != nil {
		return err
	}
	v.GameParameters = gameParameters
	v.Turns = data.Turns
	v.CurrentFaction = data.CurrentPlayer
	v.CurrentTurnStartTime = time.Unix(data.CurrentTurnStartTime, 0)
	v.GameID = data.GameID
	v.HistoryStartTurn = data.HistoryStartTurn
	return nil
}

type GameParameters struct {
	Difficulty           string
	Players              []Player
	VictoryTypes         []string
	IsOnlineMultiplayer  bool
	MultiplayerServerURL *url.URL
	BaseRuleset          string
}

func (gp *GameParameters) UnmarshalJSON(data []byte) error {
	// unmarshal to intermediate first
	var intermediate gameParametersIntermediate
	err := json.Unmarshal(data, &intermediate)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to intermediate: %w", err)
	}

	err = intermediate.Convert(gp)
	if err != nil {
		return fmt.Errorf("failed to convert intermediate: %w", err)
	}
	return nil
}

type gameParametersIntermediate struct {
	Difficulty           string   `json:"difficulty"`
	Players              []Player `json:"players"`
	VictoryTypes         []string `json:"victoryTypes"`
	IsOnlineMultiplayer  bool     `json:"isOnlineMultiplayer"`
	MultiplayerServerURL string   `json:"multiplayerServerUrl"`
	BaseRuleset          string   `json:"baseRuleset"`
}

func (data gameParametersIntermediate) Convert(v *GameParameters) error {
	parsedURL, err := url.Parse(data.MultiplayerServerURL)
	if err != nil {
		return err
	}
	v.Difficulty = data.Difficulty
	v.Players = data.Players
	v.VictoryTypes = data.VictoryTypes
	v.IsOnlineMultiplayer = data.IsOnlineMultiplayer
	v.MultiplayerServerURL = parsedURL
	v.BaseRuleset = data.BaseRuleset
	return nil
}

type Player struct {
	ChosenCiv  string `json:"chosenCiv"`
	PlayerType string `json:"playerType"`
	PlayerID   string `json:"playerId"`
}

func newSaveFromData(data saveData) (Save, error) {
	return &save{
		data: data,
	}, nil
}

// it is the caller's responsibility to close the response body
func newSaveFromResponse(resp *http.Response) (Save, error) {
	gzreader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	defer gzreader.Close()
	var data saveData
	dec := json.NewDecoder(gzreader)
	if err := dec.Decode(&data); err != nil {
		return nil, err
	}
	return newSaveFromData(data)
}

// DownloadSave to a temp file and return a Save. It is the caller's
// responsibility to Close the save.
func (server *uncivServer) DownloadSave(gameID string) (Save, error) {
	u := server.URL.JoinPath(SERVER_FILES_ROUTE, gameID)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := server.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: status code %d", ErrBadStatusCode, resp.StatusCode)
	}

	return newSaveFromResponse(resp)
}
