package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestMinOperationsBinaryArray(t *testing.T) {
	//t.Run("t1", func(t *testing.T) {
	//	assert.Equal(t, -1, leetcode.MinOperationsBinaryArray([]int{0}))
	//})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 1, leetcode.MinOperationsBinaryArray([]int{0, 0, 0}))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, -1, leetcode.MinOperationsBinaryArray([]int{0, 0, 1}))
	})

	t.Run("t4", func(t *testing.T) {
		assert.Equal(t, 3, leetcode.MinOperationsBinaryArray([]int{0, 1, 1, 1, 0, 0}))
	})

	t.Run("t5", func(t *testing.T) {
		assert.Equal(t, -1, leetcode.MinOperationsBinaryArray([]int{0, 1, 1, 1}))
	})

	t.Run("t6", func(t *testing.T) {
		assert.Equal(t, 0, leetcode.MinOperationsBinaryArray([]int{1, 1, 1}))
	})
}

func TestMinOperationsBinaryArrayV2(t *testing.T) {
	//t.Run("t1", func(t *testing.T) {
	//	assert.Equal(t, -1, leetcode.MinOperationsBinaryArrayV2([]int{0}))
	//})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 1, leetcode.MinOperationsBinaryArrayV2([]int{0, 0, 0}))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, -1, leetcode.MinOperationsBinaryArrayV2([]int{0, 0, 1}))
	})

	t.Run("t4", func(t *testing.T) {
		assert.Equal(t, 3, leetcode.MinOperationsBinaryArrayV2([]int{0, 1, 1, 1, 0, 0}))
	})

	t.Run("t5", func(t *testing.T) {
		assert.Equal(t, -1, leetcode.MinOperationsBinaryArrayV2([]int{0, 1, 1, 1}))
	})

	t.Run("t6", func(t *testing.T) {
		assert.Equal(t, 0, leetcode.MinOperationsBinaryArrayV2([]int{1, 1, 1}))
	})
}
