package models

//Player is base struct
type Player struct {
	Name   string
	Height int
	Loc    int
	Cont   bool
	Health int
}

var currentPlayer Player

//GetCurrentPlayer returns current player struct
func GetCurrentPlayer() Player {
	return currentPlayer
}

//UpdateCurrentPlayer updates the current player struct
func UpdateCurrentPlayer(np Player) {
	currentPlayer = np
}
