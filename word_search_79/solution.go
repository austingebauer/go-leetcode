package word_search_79

import (
	"strings"
)

// Time: O(N * 4^(L (len of word)))
// Space: O(L) recursion stack at most L
func exist(board [][]byte, word string) bool {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			if word[0] == board[r][c] && explore(board, word, 0, r, c) {
				return true
			}
		}
	}

	return false
}

func explore(board [][]byte, word string, step, r, c int) bool {
	// if the step is equal to the length of word, then word exists
	if step == len(word) {
		return true
	}

	// if r or c are out of bounds
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
		return false
	}

	// if the character at the step is not the next character in word
	if word[step] != board[r][c] {
		return false
	}

	// if the character has been seen before
	if strings.HasPrefix(string(board[r][c]), "_") {
		return false
	}

	// choose to mark the character as seen
	original := board[r][c]
	board[r][c] = '_' + original

	// explore each direction in search for word
	// short circuit recursion when we've found word
	exists := explore(board, word, step+1, r-1, c) ||
		explore(board, word, step+1, r+1, c) ||
		explore(board, word, step+1, r, c-1) ||
		explore(board, word, step+1, r, c+1)

	// un-choose the character
	board[r][c] = original

	return exists
}
