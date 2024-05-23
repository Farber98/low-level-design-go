package board

import "fmt"

// Board component that tracks the board
type Board struct {
	grid [][]byte
	// Strategy to check the board
	checkingStrategy CheckingStrategy
	// Precalculate data so we detect winners in constant time
	// Maps symbol to row, col and diag count
	rowCount        map[byte]map[int]int
	colCount        map[byte]map[int]int
	firstDiagCount  map[byte]int
	secondDiagCount map[byte]int
}

// Define newBoard that returns the initialized new board
type BoardInterface interface {
	// Exposed
	PrintBoard()
	PlaceSymbol(row, col int, symbol byte)
	CheckBoard(row, col int, symbol byte) bool

	// Internal
	isValidSymbolPlacement(row, col int) bool
	makeWinnerRow(row int)
	makeWinnerCol(col int)
	makeWinnerFirstDiag(row, col int)
	makeWinnerSecondDiag(row, col int)
}

func NewBoard(strategy CheckingStrategy) Board {
	// Init grid that will be initialized with 0 for bytes
	grid := make([][]byte, 3)

	for i := range grid {
		grid[i] = make([]byte, 3)
	}

	// Initialize outer maps
	rowCount := make(map[byte]map[int]int, 2)
	colCount := make(map[byte]map[int]int, 2)

	// Initialize inner maps
	rowCount['X'] = make(map[int]int, 3)
	rowCount['O'] = make(map[int]int, 3)
	colCount['X'] = make(map[int]int, 3)
	colCount['O'] = make(map[int]int, 3)

	return Board{
		grid,
		strategy,
		rowCount,
		colCount,
		make(map[byte]int, 3),
		make(map[byte]int, 3),
	}
}

func (b *Board) CheckBoard(row, col int, symbol byte) bool {
	return b.checkingStrategy.CheckBoard(b, row, col, symbol)
}

func (b *Board) PrintBoard() {
	for i := range b.grid {
		fmt.Print("|")
		for j := range b.grid[i] {
			fmt.Print(" ")
			if b.grid[i][j] != 0 {
				fmt.Print(string(b.grid[i][j]))
			} else {
				fmt.Print("_")
			}
		}
		fmt.Print(" |")
		fmt.Println()
	}
}

func (b *Board) isValidSymbolPlacement(row, col int) bool {
	if b.grid[row][col] != 0 {
		return false
	}

	if row > 2 || row < 0 || col > 2 || col < 0 {
		return false
	}

	return true
}

func (b *Board) PlaceSymbol(row, col int, symbol byte) bool {
	if !b.isValidSymbolPlacement(row, col) {
		return false
	}

	b.grid[row][col] = symbol

	// Update counts so we are able to solve in constant time
	b.rowCount[symbol][row]++
	b.colCount[symbol][col]++

	if row == col {
		b.firstDiagCount[symbol]++
	}

	if (row == 0 && col == 2) || (row == 1 && col == 1) || (row == 2 && col == 0) {
		b.secondDiagCount[symbol]++
	}

	return true
}

func (b *Board) makeWinnerRow(row int) {
	for i := range b.grid[row] {
		b.grid[row][i] = 'W'
	}
}

func (b *Board) makeWinnerCol(col int) {
	for row := 0; row < 3; row++ {
		b.grid[row][col] = 'W'
	}
}

func (b *Board) makeWinnerFirstDiag(row, col int) {
	for i := 0; i < 3; i++ {
		b.grid[i][i] = 'W'
	}
}

func (b *Board) makeWinnerSecondDiag(row, col int) {
	for i := 0; i < 3; i++ {
		b.grid[i][2-i] = 'W'
	}
}
