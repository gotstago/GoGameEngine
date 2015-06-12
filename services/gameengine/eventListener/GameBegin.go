package eventListener

import (
	//"strings"
	"github.com/gotstago/GoGameEngine/model/game"
	"log"
)

/*
This lexer function starts everything off. It determines if we are
beginning with a key/value assignment or a section.
*/
func GameBegin(listener *EventListener) ListenerFn {
	//for {
	select {
	case event := <-listener.Input:
		log.Println("got me some events...", event)
		listener.Actions <- game.GameAction{Type: game.GAME_ACTION_ANNOUNCE_BELLA, Value: "bella"}
		return GameBegin
		//			return token
	default:
		return GameBegin
		//			this.State = this.State(this) //this invokes the LexBegin method
	}
	//}
	return GameBegin
	//panic("Lexer.NextToken reached an invalid state!!")

	//return GameDealingOne
}
