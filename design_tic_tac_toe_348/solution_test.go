package design_tic_tac_toe_348

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTicTacToeXWinsRow(t *testing.T) {
	toe := Constructor(3)
	assert.Equal(t, 0, toe.Move(0, 0, 1))
	assert.Equal(t, 0, toe.Move(0, 2, 2))
	assert.Equal(t, 0, toe.Move(2, 2, 1))
	assert.Equal(t, 0, toe.Move(1, 1, 2))
	assert.Equal(t, 0, toe.Move(2, 0, 1))
	assert.Equal(t, 0, toe.Move(1, 0, 2))
	assert.Equal(t, 1, toe.Move(2, 1, 1))
}

func TestTicTacToeXWinsCol(t *testing.T) {
	toe := Constructor(3)
	assert.Equal(t, 0, toe.Move(0, 0, 1))
	assert.Equal(t, 0, toe.Move(0, 2, 2))
	assert.Equal(t, 0, toe.Move(1, 0, 1))
	assert.Equal(t, 0, toe.Move(1, 2, 2))
	assert.Equal(t, 1, toe.Move(2, 0, 1))
}

func TestTicTacToeOWinsRow(t *testing.T) {
	toe := Constructor(3)
	assert.Equal(t, 0, toe.Move(0, 0, 1))
	assert.Equal(t, 0, toe.Move(1, 0, 2))
	assert.Equal(t, 0, toe.Move(0, 2, 1))
	assert.Equal(t, 0, toe.Move(1, 1, 2))
	assert.Equal(t, 0, toe.Move(2, 2, 1))
	assert.Equal(t, 2, toe.Move(1, 2, 2))
}

func TestTicTacToeOWinsCol(t *testing.T) {
	toe := Constructor(3)
	assert.Equal(t, 0, toe.Move(0, 0, 1))
	assert.Equal(t, 0, toe.Move(0, 2, 2))
	assert.Equal(t, 0, toe.Move(0, 1, 1))
	assert.Equal(t, 0, toe.Move(1, 2, 2))
	assert.Equal(t, 0, toe.Move(1, 0, 1))
	assert.Equal(t, 2, toe.Move(2, 2, 2))
}

func TestTicTacToeNoWins(t *testing.T) {
	toe := Constructor(3)
	assert.Equal(t, 0, toe.Move(0, 0, 1))
	assert.Equal(t, 0, toe.Move(0, 1, 2))
	assert.Equal(t, 0, toe.Move(0, 2, 1))
	assert.Equal(t, 0, toe.Move(1, 2, 2))
	assert.Equal(t, 0, toe.Move(2, 1, 1))
	assert.Equal(t, 0, toe.Move(2, 2, 2))
	assert.Equal(t, 0, toe.Move(1, 1, 1))
	assert.Equal(t, 0, toe.Move(2, 0, 2))
	assert.Equal(t, 0, toe.Move(1, 0, 1))
}

func TestTicTacToeForwardDiagonalWin(t *testing.T) {
	toe := Constructor(3)
	assert.Equal(t, 0, toe.Move(2, 0, 1))
	assert.Equal(t, 0, toe.Move(0, 0, 2))
	assert.Equal(t, 0, toe.Move(1, 1, 1))
	assert.Equal(t, 0, toe.Move(0, 1, 2))
	assert.Equal(t, 1, toe.Move(0, 2, 1))
}

func TestTicTacToeBackwardDiagonalWin(t *testing.T) {
	toe := Constructor(3)
	assert.Equal(t, 0, toe.Move(0, 1, 1))
	assert.Equal(t, 0, toe.Move(0, 0, 2))
	assert.Equal(t, 0, toe.Move(0, 2, 1))
	assert.Equal(t, 0, toe.Move(1, 1, 2))
	assert.Equal(t, 0, toe.Move(1, 2, 1))
	assert.Equal(t, 2, toe.Move(2, 2, 2))
}
