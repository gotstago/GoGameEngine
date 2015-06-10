package gameengine

import (
	"testing"

	"github.com/gotstago/GoGameEngine/model/game"
)

func TestRecord(t *testing.T) {
	input := make(chan game.GameEvent)
	o := Record("tarabish", input)
	if len(o.Stages) == 0 {
		t.Error("Expected something, got nothing")
	}
}
