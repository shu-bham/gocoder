package majority_element_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/majority_element"
	"testing"
)

func TestMajorityElementHappy(t *testing.T) {
	nums := []int{3, 2, 3}
	element := majority_element.MajorityElement(nums)
	assert.Equal(t, 3, element)
}

func TestMajorityElementHappy2(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	element := majority_element.MajorityElement(nums)
	assert.Equal(t, 2, element)
}

func TestMajorityElementHappy3(t *testing.T) {
	nums := []int{3, 3, 3}
	element := majority_element.MajorityElement(nums)
	assert.Equal(t, 3, element)
}
