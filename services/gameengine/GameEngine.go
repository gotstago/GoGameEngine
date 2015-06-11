package gameengine

import (
	"github.com/gotstago/GoGameEngine/model/game"
	//"github.com/gotstago/GoGameEngine/services/gameengine"
	"log"
)

func Record(gameName string, input chan game.GameEvent) game.GameLog {
	output := game.GameLog{
		Stages: make([]game.GameStage, 0),
	}
	log.Println("Starting my listener and parser for file", gameName, "...")

	listener := BeginListening(gameName, input)
	log.Println("Starting my listener and parser for file", listener.Name, "...")
	return output
}
