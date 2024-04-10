package unciv

import "fmt"

type Save interface {
	GetCurrentPlayer() (Player, error)
}

type save struct {
	data saveData
}

func (save *save) GetCurrentPlayer() (Player, error) {
	for _, player := range save.data.GameParameters.Players {
		if player.ChosenCiv == save.data.CurrentFaction {
			return player, nil
		}
	}
	return Player{}, fmt.Errorf("%w with faction %s", ErrCouldNotFindPlayer, save.data.CurrentFaction)
}
