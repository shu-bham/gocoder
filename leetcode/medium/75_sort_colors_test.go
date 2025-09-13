package medium_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/medium"
	"testing"
)

func TestSortColors(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		arr := []int{2, 0, 2, 1, 1, 0}
		medium.SortColors(arr)
		assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, arr)
	})

	t.Run("t2", func(t *testing.T) {
		arr := []int{2, 0, 1}
		medium.SortColors(arr)
		assert.Equal(t, []int{0, 1, 2}, arr)
	})

	t.Run("t3", func(t *testing.T) {
		arr := []int{1, 0, 1, 1, 2, 0, 1, 2, 1}
		medium.SortColors(arr)
		assert.Equal(t, []int{0, 0, 1, 1, 1, 1, 1, 2, 2}, arr)
	})
}
