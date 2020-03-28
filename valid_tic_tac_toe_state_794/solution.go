package valid_tic_tac_toe_state_794

import (
	"strings"
)

func validTicTacToe(board []string) bool {
	x := 0
	o := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if string(board[i][j]) == "X" {
				x++
			} else if string(board[i][j]) == "O" {
				o++
			}
		}
	}

	oWin := 0
	xWin := 0
	winners := getWinner(board)
	for _, w := range winners {
		if w == "X" {
			xWin++
		} else {
			oWin++
		}
	}

	// x and o counts must be equal and x must be one greater than o
	if x != o && x != o+1 {
		return false
	}

	// if x wins, it must be on x's turn
	if xWin == 1 && x != o+1 {
		return false
	}

	// if o wins, it must be on o's turn
	if oWin == 1 && x != o {
		return false
	}

	return true
}

func getWinner(board []string) []string {
	b := strings.Join(board, "")
	win := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	winners := make([]string, 0)
	for _, w := range win {
		x := 0
		o := 0
		for _, s := range w {
			if string(b[s]) == "X" {
				x++
			} else if string(b[s]) == "O" {
				o++
			}
		}

		if x == 3 {
			winners = append(winners, "X")
		}

		if o == 3 {
			winners = append(winners, "O")
		}
	}

	return winners
}
