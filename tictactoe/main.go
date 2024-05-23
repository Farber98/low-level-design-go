package main

import (
	"github.com/Farber98/low-level-design/tictactoe/board"
	"github.com/Farber98/low-level-design/tictactoe/game"
	"github.com/Farber98/low-level-design/tictactoe/player"
)

func main() {
	// Get players
	player1 := player.NewPlayer("Juan", 'X')
	player2 := player.NewPlayer("ChainLink", 'O')

	// Pick a strategy
	strategy := &board.ConstantCheckingStrategy{}

	// Get board
	board := board.NewBoard(strategy)

	// Generate game with players and board
	game := game.NewGame(board, player1, player2)

	game.Play()
}

/*
<>

INFORMATION:
- Design a tic tac toe game

QUESTIONS & ASSUMPTIONS:
- 3x3 grid
- 2 players
- 3 possible outcomes, p1 wins, p2 wins, draw
- To win, we can use diag, reverse diag, row or col

CONSTRAINTS:
- check out of borders
- check place already filled

COMPONENTS:
- Player component that tells which symbol is
	- Define a struct that contains a name and symbol
	- Defin newPlayer that returns an initialized player

- Board component that tracks the board
	- Define a struct that contains grid [][]byte
	- Define newBoard that returns the initialized new board

- Game component that tracks the game development (checks if someone wins, draws, etc)
	- Game struct will have player 1, player 2 and board
	- It will be responsible of
		- Initializing new game by receiving players and board
		- nextTurn() functionality will alternate between player 1 and 2 turns. We can make the first turn random and next deterministic.
		- Place functionality to place a symbol on a user given x y coordinates
			- If already placed, make user retry
		- After placing, we should have a function that checks if recent player that put a symbol won or if we have a draw
			- hasWon will receive the player and scan the board to check if it's the winner, printing that he won and ending game
			- isDraw will check if the board is fully filled and we can't continue placing, ending the game


EXAMPLE:

INTUITION:

APPROACH:

TIME COMPLEXITY:

SPACE COMPLEXITY:

<>
*/
