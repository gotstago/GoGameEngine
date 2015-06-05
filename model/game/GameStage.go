package game

type GameStage struct {
	Name    string       `json:"name"`
	Actions []GameAction `json:"actions"`
}
