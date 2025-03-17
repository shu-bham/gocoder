package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestApplyOperations(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, []int{1, 4, 2, 0, 0, 0}, leetcode.ApplyOperations([]int{1, 2, 2, 1, 1, 0}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, []int{1, 0, 0}, leetcode.ApplyOperations([]int{0, 1, 0}))
	})

}

func TestApplyOperationsV2(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, []int{1, 4, 2, 0, 0, 0}, leetcode.ApplyOperationsV2([]int{1, 2, 2, 1, 1, 0}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, []int{1, 0, 0}, leetcode.ApplyOperationsV2([]int{0, 1, 0}))
	})

}
