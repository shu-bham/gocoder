package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestPivotArray(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		pivotArray := leetcode.PivotArray([]int{9, 12, 5, 10, 14, 3, 10}, 10)
		assert.Equal(t, []int{9, 5, 3, 10, 10, 12, 14}, pivotArray)
	})
}

func TestPivotV2Array(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		pivotArray := leetcode.PivotArrayV2([]int{9, 12, 5, 10, 14, 3, 10}, 10)
		assert.Equal(t, []int{9, 5, 3, 10, 10, 12, 14}, pivotArray)
	})
}
