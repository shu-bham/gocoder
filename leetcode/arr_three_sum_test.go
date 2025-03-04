package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	pos := leetcode.ThreeSum(nums)
	assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, pos)
}

func TestThreeSum1(t *testing.T) {
	nums := []int{0, 0, 0}
	pos := leetcode.ThreeSum(nums)
	assert.Equal(t, [][]int{{0, 0, 0}}, pos)
}
