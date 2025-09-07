package hard_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/hard"
	"testing"
)

func TestBuildMatrix(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		k := 3
		rowCond := [][]int{{1, 2}, {3, 2}}
		colCond := [][]int{{2, 1}, {3, 2}}
		res := hard.BuildMatrix(k, rowCond, colCond)
		assert.Equal(t, [][]int{{3, 0, 0}, {0, 0, 1}, {0, 2, 0}}, res)
	})

	t.Run("t2", func(t *testing.T) {
		k := 3
		rowCond := [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}}
		colCond := [][]int{{2, 1}}
		res := hard.BuildMatrix(k, rowCond, colCond)
		assert.Equal(t, [][]int{}, res)
	})
}
