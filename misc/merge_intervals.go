package misc

import "sort"

func MergeIntervals(intervals [][]int) [][]int {
	res := make([][]int, 0)

	if len(intervals) == 0 {
		return res
	}

	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		merged := res[len(res)-1]
		curr := intervals[i]

		if curr[0] <= merged[1] {
			merged[1] = max(merged[1], curr[1])
		} else {
			res = append(res, []int{curr[0], curr[1]})
		}
	}

	return res

}

func MergeIntervalsInplace(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	lastMergedIdx := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[lastMergedIdx][1] >= intervals[i][0] {
			intervals[lastMergedIdx][1] = max(intervals[lastMergedIdx][1], intervals[i][1])
		} else {
			lastMergedIdx++
			intervals[lastMergedIdx] = intervals[i]
		}
	}

	return intervals[:lastMergedIdx+1]

}
