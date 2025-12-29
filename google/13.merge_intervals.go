package google

import "sort"

func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var ans [][]int
	ans = append(ans, intervals[0])

	for i := 1; i < len(intervals); i++ {
		lastInterval := ans[len(ans)-1]
		currInterval := intervals[i]

		if currInterval[0] <= lastInterval[1] {
			lastInterval[1] = max(lastInterval[1], currInterval[1])
		} else {
			ans = append(ans, currInterval)
		}
	}

	return ans
}
