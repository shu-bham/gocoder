package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestCanJump(t *testing.T) {
	arr := []int{1, 2, 3}
	can := leetcode.CanJump(arr)
	assert.Equal(t, true, can)
}

func TestCanJump1(t *testing.T) {
	arr := []int{1}
	can := leetcode.CanJump(arr)
	assert.Equal(t, true, can)
}

func TestCanJump2(t *testing.T) {
	arr := []int{1, 0}
	can := leetcode.CanJump(arr)
	assert.Equal(t, true, can)
}

func TestCanJump3(t *testing.T) {
	arr := []int{3, 0, 0, 0}
	can := leetcode.CanJump(arr)
	assert.Equal(t, true, can)
}

func TestCanJump4(t *testing.T) {
	arr := []int{3, 2, 1, 0, 4}
	can := leetcode.CanJump(arr)
	assert.Equal(t, false, can)
}

func TestCanJump5(t *testing.T) {
	arr := []int{0}
	can := leetcode.CanJump(arr)
	assert.Equal(t, true, can)
}

func TestCanJump6(t *testing.T) {
	arr := []int{2, 3, 1, 1, 4}
	can := leetcode.CanJumpV2(arr)
	assert.Equal(t, true, can)
}
