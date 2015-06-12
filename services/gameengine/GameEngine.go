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
	stage := game.GameStage{}
	stage.Actions = make([]game.GameAction, 0)
	listener := BeginListening(gameName, input)
	//listener.State = listener.State(listener)
	action := listener.NextAction()
	stage.Actions = append(stage.Actions, action)
	output.Stages = append(output.Stages, stage)
	log.Println("Starting my listener and parser for file", listener.Name, action.Value, "...")
	return output
}
