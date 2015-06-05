package game

type GameLog struct {
	//FileName string       `json:"fileName"`
	Stages []GameStage `json:"stages"`
}
