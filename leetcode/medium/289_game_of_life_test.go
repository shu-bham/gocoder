package medium_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/medium"
	"testing"
)

func TestGameOfLifeV1(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
		medium.GameOfLifeV1(board)
		expected := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}
		assert.Equal(t, expected, board)
	})

	t.Run("t2", func(t *testing.T) {
		board := [][]int{{1, 1}, {1, 0}}
		medium.GameOfLifeV1(board)
		expected := [][]int{{1, 1}, {1, 1}}
		assert.Equal(t, expected, board)
	})
}

func TestGameOfLifeOptimised(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
		medium.GameOfLifeOptimised(board)
		expected := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}
		assert.Equal(t, expected, board)
	})

	t.Run("t2", func(t *testing.T) {
		board := [][]int{{1, 1}, {1, 0}}
		medium.GameOfLifeOptimised(board)
		expected := [][]int{{1, 1}, {1, 1}}
		assert.Equal(t, expected, board)
	})
}
