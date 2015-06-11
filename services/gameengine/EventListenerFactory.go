package gameengine

import (
	"github.com/gotstago/GoGameEngine/model/game"
	"github.com/gotstago/GoGameEngine/services/gameengine/eventListener"
)

/*
Start a new listener with a given input channel. This returns the
instance of the listener and a channel of actions. Reading this stream
is the way to parse a given input and perform processing.
*/
func BeginListening(name string, input chan game.GameEvent) *eventListener.EventListener {
	l := &eventListener.EventListener{
		Name:    name,
		Input:   input,
		State:   eventListener.GameBegin,
		Actions: make(chan game.GameAction),
	}

	return l
}
