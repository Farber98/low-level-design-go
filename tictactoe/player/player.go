package player

// Define a struct that contains a name and symbol
type Player struct {
	Name   string
	Symbol byte
}

// Player component that tells which symbol is
// Defin newPlayer that returns an initialized player
func NewPlayer(name string, symbol byte) Player {
	return Player{name, symbol}
}
