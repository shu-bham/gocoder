package medium_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/medium"
	"testing"
)

func TestMostProfitablePath(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}
		bob := 3
		amount := []int{-2, 4, 2, -4, 6}
		profit := medium.MostProfitablePath(edges, bob, amount)
		assert.Equal(t, 6, profit)
	})

	t.Run("t2", func(t *testing.T) {
		edges := [][]int{{0, 1}}
		bob := 1
		amount := []int{-7280, 2350}
		profit := medium.MostProfitablePath(edges, bob, amount)
		assert.Equal(t, -7280, profit)
	})

	t.Run("t3", func(t *testing.T) {
		edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}, {2, 5}, {5, 6}, {3, 7}, {6, 8}, {6, 9}}
		bob := 6
		amount := []int{-2, 2, 4, -2, -4, -2, 4, -2, 2, -2}
		profit := medium.MostProfitablePath(edges, bob, amount)
		assert.Equal(t, 4, profit)
	})

	t.Run("t4", func(t *testing.T) {
		edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {3, 5}, {3, 6}, {4, 7}, {4, 8}, {8, 9}}
		bob := 9
		amount := []int{4, -8, 6, 10, -2, -12, -20, -6, -16, -32}
		profit := medium.MostProfitablePath(edges, bob, amount)
		assert.Equal(t, 9, profit)
	})

}
