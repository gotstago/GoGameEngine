package game

import "fmt"

type GameAction struct {
	Type   GameActionType `json:"type"`
	Actor  string         `json:"actor"`
	Target string         `json:"target"`
	Value  string         `json:"value"`
}

func (this GameAction) String() string {
	switch this.Type {
	case GAME_ACTION_END_OF_GAME:
		return "Game Over"

	case GAME_ACTION_ERROR:
		return this.Value

	case GAME_ACTION_BID:
		return fmt.Sprintf("%s bids %s", this.Actor, this.Value)
	}

	return fmt.Sprintf("%q", this.Value)
}
