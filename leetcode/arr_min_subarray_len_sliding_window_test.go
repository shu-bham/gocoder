package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	l := leetcode.MinSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	assert.Equal(t, 2, l)
}

func TestMinSubArrayLen1(t *testing.T) {
	l := leetcode.MinSubArrayLen(7, []int{1, 1, 1})
	assert.Equal(t, 0, l)
}

func TestMinSubArrayLen2(t *testing.T) {
	l := leetcode.MinSubArrayLen(6, []int{2, 3, 1})
	assert.Equal(t, 3, l)
}
