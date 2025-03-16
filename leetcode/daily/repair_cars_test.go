package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestRepairCars(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, int64(16), leetcode.RepairCars([]int{4, 2, 3, 1}, 10))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, int64(16), leetcode.RepairCars([]int{5, 1, 8}, 6))
	})
}
