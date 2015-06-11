// EventListener.go
package eventListener

import (
	//"fmt"
	"github.com/gotstago/GoGameEngine/model/game"
)

type EventListener struct {
	Name    string
	Input   chan game.GameEvent
	Actions chan game.GameAction
	State   ListenerFn

	//	Start int
	//	Pos   int
	//	Width int
}
