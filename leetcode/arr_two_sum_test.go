package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestTwoSumHappy(t *testing.T) {
	nums := []int{3, 2, 3}
	arr := leetcode.TwoSum(nums, 6)
	assert.Equal(t, []int{0, 2}, arr)
}

func TestTwoSumHappy2(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	arr := leetcode.TwoSum(nums, 9)
	assert.Equal(t, []int{0, 1}, arr)
}
