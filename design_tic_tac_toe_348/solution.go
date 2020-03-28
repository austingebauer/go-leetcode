package design_tic_tac_toe_348

import "math"

type TicTacToe struct {
	rowCounts         []int
	colCounts         []int
	forwardDiagCount  int
	backwardDiagCount int
}

/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	return TicTacToe{
		rowCounts: make([]int, n),
		colCounts: make([]int, n),
	}
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
	// Note: intuition is that each row, col, and diagonal has a count.
	// For a player to win, it's count must be equal to the row, col, or diagonal length.
	// Let's choose +1 for player 1 and -1 from player 2. If the absolute value of the
	// count for row, col, diagonals ever reaches the size of the grid, then the player
	// has won. This behavior assumes each move is valid and into an empty block.

	// the count for x (player 1) is 1
	// the count for o (player 2) is -1
	count := 1
	if player == 2 {
		count = -1
	}

	// increase the counts for each row and column
	this.rowCounts[row] += count
	this.colCounts[col] += count

	// increase backward slash diagonal count
	if row == col {
		this.backwardDiagCount += count
	}

	// increase forward slash diagonal count
	// tricky observation for the forward slash diagonal
	// works for matching on columns as well
	if row == len(this.rowCounts)-col-1 {
		this.forwardDiagCount += count
	}

	// if the count at this move becomes 3 or -3, the current player
	// just won the game.
	boardSize := len(this.rowCounts)
	if int(math.Abs(float64(this.rowCounts[row]))) == boardSize ||
		int(math.Abs(float64(this.colCounts[col]))) == boardSize ||
		int(math.Abs(float64(this.forwardDiagCount))) == boardSize ||
		int(math.Abs(float64(this.backwardDiagCount))) == boardSize {
		return player
	}

	// no-one won the game
	return 0
}
