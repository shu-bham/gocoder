package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestMinOperationsToMakeUniVal(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 4, leetcode.MinOperationsToMakeUniVal([][]int{{2, 4}, {6, 8}}, 2))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 5, leetcode.MinOperationsToMakeUniVal([][]int{{1, 5}, {2, 3}}, 1))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, -1, leetcode.MinOperationsToMakeUniVal([][]int{{1, 2}, {3, 4}}, 2))
	})

	t.Run("t4", func(t *testing.T) {
		assert.Equal(t, 25, leetcode.MinOperationsToMakeUniVal([][]int{{529, 529, 989}, {989, 529, 345}, {989, 805, 69}}, 92))
	})
}
