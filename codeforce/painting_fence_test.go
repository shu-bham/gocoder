package codeforce_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/codeforce"
	"testing"
)

func TestMinNoOfStrokes(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 1, codeforce.MinNoOfStrokes([]int{5}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 3, codeforce.MinNoOfStrokes([]int{5, 3, 3}))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, 3, codeforce.MinNoOfStrokes([]int{2, 2, 1, 2, 1}))
	})

	t.Run("t4", func(t *testing.T) {
		assert.Equal(t, 1, codeforce.MinNoOfStrokes([]int{5}))
	})

	t.Run("t5", func(t *testing.T) {
		assert.Equal(t, 6, codeforce.MinNoOfStrokes([]int{3, 2, 4, 1, 4, 2, 1}))
	})

	t.Run("t6", func(t *testing.T) {
		assert.Equal(t, 3, codeforce.MinNoOfStrokes([]int{3, 1, 2, 2, 2}))
	})

	t.Run("t7 - Single fence with zero height", func(t *testing.T) {
		assert.Equal(t, 0, codeforce.MinNoOfStrokes([]int{0}))
	})

	t.Run("t8 - All fences same height", func(t *testing.T) {
		assert.Equal(t, 5, codeforce.MinNoOfStrokes([]int{5, 5, 5, 5, 5}))
	})

	t.Run("t9 - Strictly increasing heights", func(t *testing.T) {
		assert.Equal(t, 5, codeforce.MinNoOfStrokes([]int{1, 2, 3, 4, 5}))
	})

	t.Run("t10 - Strictly decreasing heights", func(t *testing.T) {
		assert.Equal(t, 5, codeforce.MinNoOfStrokes([]int{5, 4, 3, 2, 1}))
	})

	t.Run("t11 - All fences have height 1", func(t *testing.T) {
		assert.Equal(t, 1, codeforce.MinNoOfStrokes([]int{1, 1, 1, 1, 1}))
	})

	t.Run("t12 - Large input case", func(t *testing.T) {
		largeInput := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			largeInput[i] = i + 1
		}
		assert.Equal(t, 1000, codeforce.MinNoOfStrokes(largeInput))
	})

	t.Run("t13 - Mixed heights with peaks and valleys", func(t *testing.T) {
		assert.Equal(t, 9, codeforce.MinNoOfStrokes([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}))
	})
}
