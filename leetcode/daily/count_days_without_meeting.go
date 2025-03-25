package leetcode

import (
	"slices"
	"sort"
)

// Medium: https://leetcode.com/problems/count-days-without-meetings
func CountDaysWithoutMeeting(days int, meetings [][]int) int {

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	lastMergedIndex := 0

	for i := 1; i < len(meetings); i++ {
		merged := meetings[lastMergedIndex]
		curr := meetings[i]

		if curr[0] <= merged[1] {
			merged[1] = max(merged[1], curr[1])
		} else {
			lastMergedIndex++
			meetings[lastMergedIndex] = meetings[i]
		}
	}

	res := days
	for i := 0; i <= lastMergedIndex; i++ {
		daysInMeet := meetings[i][1] - meetings[i][0] + 1
		res -= daysInMeet
	}

	return res

}

// optimised
func CountDaysWithoutMeetingV2(days int, meetings [][]int) int {
	slices.SortFunc(meetings, func(a, b []int) int {
		return a[0] - b[0]
	})
	ans, lastEndDay := 0, 0
	for _, meeting := range meetings {
		if meeting[0] > lastEndDay {
			ans += meeting[0] - lastEndDay - 1
		}
		lastEndDay = max(lastEndDay, meeting[1])
		if meeting[1] >= days {
			break
		}
	}
	if days > lastEndDay {
		ans += days - lastEndDay
	}
	return ans
}
