package gameengine

import (
	"testing"

	"github.com/gotstago/GoGameEngine/model/game"
)

func TestRecordHasAtLeastOneStage(t *testing.T) {
	input := make(chan game.GameEvent)
	o := Record("tarabish", input)
	//input <- game.GameEvent{Type: game.GAME_EVENT_ANNOUNCE_BELLA, Value: ""}
	if len(o.Stages) != 0 {
		t.Error("Expected something, got nothing")
	}
}
