package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestDivideArray(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, false, leetcode.DivideArray([]int{1, 2, 3, 4}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, true, leetcode.DivideArray([]int{1, 1, 3, 3}))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, true, leetcode.DivideArray([]int{3, 2, 3, 2, 2, 2}))
	})

}
