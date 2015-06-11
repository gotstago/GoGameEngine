package eventListener

import (
	//"strings"

	//"github.com/gotstago/GoGameEngine/services/lexer/lexertoken"
)

/*
This lexer function starts everything off. It determines if we are
beginning with a key/value assignment or a section.
*/
func GameBidding(listener *EventListener) ListenerFn {
	return GameDealingTwo
}
