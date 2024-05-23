package game

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Farber98/low-level-design/tictactoe/board"
	"github.com/Farber98/low-level-design/tictactoe/player"
)

// Game component that tracks the game development (checks if someone wins, draws, etc)

// Game struct will have player 1, player 2 and board
type Game struct {
	Board            board.Board
	Player1, Player2 player.Player
	currTurn         player.Player
	turns            int
}

// Initializing new game by receiving players and board
func NewGame(board board.Board, player1, player2 player.Player) Game {
	return Game{board, player1, player2, player1, 1}
}

type GameInterface interface {
	// Exposed
	Play() // Will be the function that starts the game and iterates until someone wins or we have a draw

	// Internal
	swapTurn()          // Will be the function that defines whose next turn is
	place(row, col int) // Places player symbol in given row col coordinates. If already placed retry.
	hasWon() bool       // Checks if the player that recently placed a val, has won
	isDraw()            // Checks if board has been filled and nobody won
}

func (g *Game) Play() {
	// Repeats the logic until someone has won or is draw
	for {
		// Print board
		g.Board.PrintBoard()

		fmt.Printf("TURN: %v\n", g.currTurn.Name)
		// Grab input
		// Input for row
		// Input for col
		row := readInput("Pick row:")
		col := readInput("Pick col:")
		if row < 0 || col < 0 || row > 2 || col > 2 {
			continue
		}

		// If placement is invalid, retries
		if !g.Board.PlaceSymbol(row, col, g.currTurn.Symbol) {
			continue
		}

		if g.hasWon(row, col) {
			fmt.Println(g.currTurn.Name + " " + "WON")
			g.Board.PrintBoard()
			break
		}

		if g.isDraw() {
			fmt.Println("DRAW")
			break
		}

		g.swapTurn()
	}
}
func (g *Game) swapTurn() {
	if g.currTurn == g.Player1 {
		g.currTurn = g.Player2
	} else {
		g.currTurn = g.Player1
	}
}

func (g *Game) hasWon(row, col int) bool {
	return g.Board.CheckBoard(row, col, g.currTurn.Symbol)
}

func (g *Game) isDraw() bool {
	// Check if our board is full
	return g.turns >= 9
}

func readInput(input string) int {
	fmt.Print(input + " ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return int(scanner.Bytes()[0] - '0')
}
