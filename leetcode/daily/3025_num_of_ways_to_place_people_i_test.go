package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestNumberOfPairs(t *testing.T) {

	t.Run("tc1", func(t *testing.T) {
		points := [][]int{{1, 1}, {2, 2}, {3, 3}}
		pairs := leetcode.NumberOfPairs(points)
		assert.Equal(t, 0, pairs)
	})

	t.Run("tc2", func(t *testing.T) {
		points := [][]int{{6, 2}, {4, 4}, {2, 6}}
		pairs := leetcode.NumberOfPairs(points)
		assert.Equal(t, 2, pairs)
	})

	t.Run("tc3", func(t *testing.T) {
		points := [][]int{{3, 1}, {1, 3}, {1, 1}}
		pairs := leetcode.NumberOfPairs(points)
		assert.Equal(t, 2, pairs)
	})

	t.Run("tc4", func(t *testing.T) {
		points := [][]int{{0, 1}, {5, 2}}
		pairs := leetcode.NumberOfPairs(points)
		assert.Equal(t, 0, pairs)
	})

}
