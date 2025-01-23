package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestMajorityElementHappy(t *testing.T) {
	nums := []int{3, 2, 3}
	element := leetcode.MajorityElement(nums)
	assert.Equal(t, 3, element)
}

func TestMajorityElementHappy2(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	element := leetcode.MajorityElement(nums)
	assert.Equal(t, 2, element)
}

func TestMajorityElementHappy3(t *testing.T) {
	nums := []int{3, 3, 3}
	element := leetcode.MajorityElement(nums)
	assert.Equal(t, 3, element)
}
