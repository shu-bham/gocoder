package misc_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/misc"
	"testing"
)

func TestMergeIntervals(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, [][]int{{1, 2}, {3, 4}}, misc.MergeIntervals([][]int{{1, 2}, {3, 4}}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, [][]int{{1, 4}}, misc.MergeIntervals([][]int{{1, 3}, {3, 4}}))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, [][]int{{1, 2}, {3, 6}}, misc.MergeIntervals([][]int{{1, 2}, {3, 6}, {5, 6}}))
	})

	t.Run("t4", func(t *testing.T) {
		assert.Equal(t, [][]int{{1, 6}}, misc.MergeIntervals([][]int{{1, 2}, {3, 6}, {5, 6}, {1, 3}}))
	})
}
