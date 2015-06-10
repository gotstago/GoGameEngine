package gameengine

import (
	"github.com/gotstago/GoGameEngine/model/game"
	//"../../model/game"
)

func Record(gameName string, input chan game.GameEvent) game.GameLog {
	output := game.GameLog{
		Stages: make([]game.GameStage, 0),
	}
	return output
}
