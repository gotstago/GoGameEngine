// EventListener.go
package eventListener

import (
	"github.com/gotstago/GoGameEngine/model/game"
	"log"
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

/*
Return the next token from the channel
*/
func (this *EventListener) NextAction() game.GameAction {
	for {
		select {
		case action := <-this.Actions:
			log.Println("got me some actions...", action.Type)
			return action
		default:
			log.Println("no actions yet...")
			this.State = this.State(this) //this invokes the LexBegin method
			log.Println("lets go another round...")
		}
	}

	panic("Lexer.NextToken reached an invalid state!!")
}
