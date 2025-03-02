package leetcode

import (
	"slices"
)

// Medium : https://leetcode.com/problems/h-index/description/
func HIndex(citations []int) int {
	slices.Sort(citations)
	n := len(citations)
	p := citations[n-1]
	temp := make([]int, p+1)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= citations[i]; j++ {
			temp[j]++
		}
	}

	for i := len(temp) - 1; i >= 0; i-- {
		if temp[i] >= i {
			return i
		}
	}

	return -1
}
