package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestMinimumCostWalkInWeightedGraph(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, []int{1, -1}, leetcode.MinimumCostWalkInWeightedGraph(5, [][]int{{0, 1, 7}, {1, 3, 7}, {1, 2, 1}}, [][]int{{0, 3}, {3, 4}}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, []int{0}, leetcode.MinimumCostWalkInWeightedGraph(3, [][]int{{0, 2, 7}, {0, 1, 15}, {1, 2, 6}, {1, 2, 1}}, [][]int{{1, 2}}))
	})
}
