package board

type CheckingStrategy interface {
	CheckBoard(board *Board, row, col int, symbol byte) bool
	checkRows(board *Board, row int, symbol byte) bool
	checkCols(board *Board, col int, symbol byte) bool
	checkDiags(board *Board, row, col int, symbol byte) bool
	checkFirstDiag(board *Board, row, col int, symbol byte) bool
	checkSecondDiag(board *Board, row, col int, symbol byte) bool
}

type ConstantCheckingStrategy struct{}

func (s *ConstantCheckingStrategy) CheckBoard(board *Board, row, col int, symbol byte) bool {
	return s.checkRows(board, row, symbol) ||
		s.checkCols(board, col, symbol) ||
		s.checkDiags(board, row, col, symbol)
}

func (s *ConstantCheckingStrategy) checkRows(board *Board, row int, symbol byte) bool {
	if board.rowCount[symbol][row] == 3 {
		board.makeWinnerRow(row)
		return true
	}

	return false
}

func (s *ConstantCheckingStrategy) checkCols(board *Board, col int, symbol byte) bool {
	if board.colCount[symbol][col] == 3 {
		board.makeWinnerCol(col)
		return true
	}

	return false
}

func (s *ConstantCheckingStrategy) checkDiags(board *Board, row, col int, symbol byte) bool {
	return s.checkFirstDiag(board, row, col, symbol) || s.checkSecondDiag(board, row, col, symbol)
}

func (s *ConstantCheckingStrategy) checkFirstDiag(board *Board, row, col int, symbol byte) bool {
	if board.firstDiagCount[symbol] == 3 {
		board.makeWinnerFirstDiag(row, col)
		return true
	}

	return false
}

func (s *ConstantCheckingStrategy) checkSecondDiag(board *Board, row, col int, symbol byte) bool {
	if board.secondDiagCount[symbol] == 3 {
		board.makeWinnerSecondDiag(row, col)
		return true
	}

	return false
}

type LinearCheckingStrategy struct{}

func (l *LinearCheckingStrategy) CheckBoard(board *Board, row, col int, symbol byte) bool {
	return l.checkRows(board, row, symbol) ||
		l.checkCols(board, col, symbol) ||
		l.checkDiags(board, row, col, symbol)
}

func (l *LinearCheckingStrategy) checkRows(board *Board, row int, symbol byte) bool {
	for col := range board.grid[row] {
		if board.grid[row][col] != symbol {
			return false
		}
	}

	board.makeWinnerRow(row)
	return true
}

func (l *LinearCheckingStrategy) checkCols(board *Board, col int, symbol byte) bool {
	for row := 0; row < 3; row++ {
		if board.grid[row][col] != symbol {
			return false
		}
	}

	board.makeWinnerCol(col)
	return true
}

func (l *LinearCheckingStrategy) checkDiags(board *Board, row, col int, symbol byte) bool {
	// Check if something was placed in first diag
	return l.checkFirstDiag(board, row, col, symbol) || l.checkSecondDiag(board, row, col, symbol)
}

func (l *LinearCheckingStrategy) checkFirstDiag(board *Board, row, col int, symbol byte) bool {
	if row != col {
		return false
	}

	for i := 0; i < 3; i++ {
		if board.grid[i][i] != symbol {
			return false
		}
	}

	board.makeWinnerFirstDiag(row, col)
	return true
}

func (l *LinearCheckingStrategy) checkSecondDiag(board *Board, row, col int, symbol byte) bool {
	if (row != 0 || col != 2) && (row != 1 || col != 1) && (row != 2 || col != 0) {
		return false
	}

	for i := 0; i < 3; i++ {
		if board.grid[i][2-i] != symbol {
			return false
		}
	}

	board.makeWinnerSecondDiag(row, col)
	return true
}
