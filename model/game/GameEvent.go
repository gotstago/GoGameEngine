package game

import "fmt"

type GameEvent struct {
	Type   GameEventType `json:"type"`
	Actor  string        `json:"actor"`
	Target string        `json:"target"`
	Value  string        `json:"value"`
}

func (this GameEvent) String() string {
	switch this.Type {

	case GAME_EVENT_ERROR:
		return this.Value

	case GAME_EVENT_BID:
		return fmt.Sprintf("%s bids %s", this.Actor, this.Value)
	}

	return fmt.Sprintf("%q", this.Value)
}
