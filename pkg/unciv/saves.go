package unciv

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type Save interface {
}

type save struct {
	data saveData
}

type saveData struct {
	GameParameters       GameParameters `json:"gameParameters"`
	Turns                int            `json:"turns"`
	CurrentPlayer        string         `json:"currentPlayer"`
	CurrentTurnStartTime time.Time      `json:"currentTurnStartTime,int"`
	GameID               string         `json:"gameId"`
	HistoryStartTurn     int            `json:"historyStartTurn"`
}

type GameParameters struct {
	Difficulty           string   `json:"difficulty"`
	Players              []Player `json:"players"`
	VictoryTypes         []string `json:"victoryTypes"`
	IsOnlineMultiplayer  bool     `json:"isOnlineMultiplayer"`
	MultiplayerServerURL *url.URL `json:"multiplayerServerUrl,string"`
	BaseRuleset          string   `json:"baseRuleset"`
}

func (gp *GameParameters) UnmarshalJSON(data []byte) error {
	// unmarshal into a simple dict
	var rawObj map[string]json.RawMessage
	err := json.Unmarshal(data, &rawObj)
	if err != nil {
		return fmt.Errorf("failed to convert to raw object: %w", err)
	}

	/* This is one of the worst things I've ever made in this language.
	   Here be dragons.
	   You have been warned.
	*/
	for k, v := range rawObj {
		switch k {
		case "multiplayerServerUrl":
			gp.MultiplayerServerURL, err = url.Parse(strings.Trim(string(v), "\""))
			if err != nil {
				return err
			}
		default:
			for _, field := range reflect.VisibleFields(reflect.TypeOf(*gp)) {
				tag, ok := field.Tag.Lookup("json")
				if !ok {
					tag = field.Name
				}
				key := strings.Split(tag, ",")[0]
				if key != k {
					continue
				}
				dest := reflect.New(field.Type).Elem()
				destPointer := dest.Addr().Interface()
				err := json.Unmarshal(v, destPointer)
				if err != nil {
					return fmt.Errorf("failed to unmarshal into reflect: %w", err)
				}
				reflect.ValueOf(gp).Elem().FieldByName(field.Name).Set(dest)
			}
		}
	}

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
	u := server.URL.JoinPath(SERVER_FILES_ROUTE)
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
