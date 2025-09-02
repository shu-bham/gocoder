package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestMaxPointsFromGridQueries(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		grid := [][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}}
		queries := []int{5, 6, 2}
		ints := leetcode.MaxPointsFromGridQueries(grid, queries)
		assert.Equal(t, []int{5, 8, 1}, ints)
	})

	t.Run("t2", func(t *testing.T) {
		grid := [][]int{{5, 2, 1}, {1, 1, 2}}
		queries := []int{3}
		ints := leetcode.MaxPointsFromGridQueries(grid, queries)
		assert.Equal(t, []int{0}, ints)
	})
}
