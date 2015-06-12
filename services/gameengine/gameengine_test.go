package gameengine

import (
	"testing"

	"github.com/gotstago/GoGameEngine/model/game"
)

func TestRecordHasAtLeastOneStage(t *testing.T) {
	input := make(chan game.GameEvent, 3)
	input <- game.GameEvent{Type: game.GAME_EVENT_ANNOUNCE_BELLA, Value: ""}

	o := Record("tarabish", input)

	if len(o.Stages) == 0 {
		t.Error("Expected something, got nothing")
	}
}
