package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestCountDaysWithoutMeeting(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 2, leetcode.CountDaysWithoutMeeting(10, [][]int{{5, 7}, {1, 3}, {9, 10}}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 1, leetcode.CountDaysWithoutMeeting(5, [][]int{{2, 4}, {1, 3}}))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, 0, leetcode.CountDaysWithoutMeeting(6, [][]int{{1, 6}}))
	})
}

func TestCountDaysWithoutMeetingV2(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 2, leetcode.CountDaysWithoutMeetingV2(10, [][]int{{5, 7}, {1, 3}, {9, 10}}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 1, leetcode.CountDaysWithoutMeetingV2(5, [][]int{{2, 4}, {1, 3}}))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, 0, leetcode.CountDaysWithoutMeetingV2(6, [][]int{{1, 6}}))
	})
}
