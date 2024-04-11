package main

import (
	"flag"

	"github.com/joshbarrass/UncivNotifyBot/pkg/unciv"
	"github.com/sirupsen/logrus"
)

func main() {
	server := unciv.NewDefaultUncivServer()
	flag.Parse()
	gameID := flag.Args()[0]
	save, err := server.DownloadSave(gameID)
	if err != nil {
		logrus.Panicf("Failed to download save: %s", err)
	}
	currentPlayer, err := save.GetCurrentPlayer()
	if err != nil {
		logrus.Panicf("Failed to get current player: %s", err)
	}
	logrus.Infof("Current player: %s (%s)", currentPlayer.ChosenCiv, currentPlayer.PlayerID)
}
