package medium_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/medium"
	"testing"
)

func TestMinOperations(t *testing.T) {

	t.Run("tc1", func(t *testing.T) {
		assert.Equal(t, []int{1, 1, 3}, medium.MinOperations("110"))
	})

	t.Run("tc2", func(t *testing.T) {
		assert.Equal(t, []int{11, 8, 5, 4, 3, 4}, medium.MinOperations("001011"))
	})
}

func TestMinOperationsV2(t *testing.T) {

	t.Run("tc1", func(t *testing.T) {
		assert.Equal(t, []int{1, 1, 3}, medium.MinOperationsV2("110"))
	})

	t.Run("tc2", func(t *testing.T) {
		assert.Equal(t, []int{11, 8, 5, 4, 3, 4}, medium.MinOperationsV2("001011"))
	})
}
