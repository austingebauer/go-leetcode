package game_of_life_289

func gameOfLife(board [][]int) {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			cs := board[r][c]
			liveNeighbors := getLiveNeighbors(board, r, c, len(board), len(board[r]))

			if cs == 0 && liveNeighbors == 3 {
				// Rule 4: Was a 0 (dead) and going to 1 (alive)
				board[r][c] = -2
			} else if cs == 1 && (liveNeighbors < 2 || liveNeighbors > 3) {
				// Rule 1,3: Was a 1 (live) and going to 0 (dead)
				board[r][c] = -1
			}
		}
	}

	// Swap -1 (1 -> 0) and -2 (0 -> 1) in place
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			if board[r][c] == -1 {
				board[r][c] = 0
			} else if board[r][c] == -2 {
				board[r][c] = 1
			}
		}
	}
}

func getLiveNeighbors(board [][]int, r, c, rows, cols int) int {
	cnt := 0
	for cRow := r - 1; cRow <= r+1; cRow++ {
		if cRow > -1 && cRow < rows {
			for cCol := c - 1; cCol <= c+1; cCol++ {
				if cCol > -1 && cCol < cols {
					if (cCol != c || cRow != r) &&
						// Remember that previously live and going to die cells are -1
						(board[cRow][cCol] == -1 || board[cRow][cCol] == 1) {
						cnt++
					}
				}
			}
		}
	}

	return cnt
}
