package google

import (
	"math"
	"sort"
)

// dp
func MaxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	if len(envelopes) == 1 {
		return 1
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	lis := make([]int, len(envelopes))
	for i := range lis {
		lis[i] = 1
	}

	ans := 0
	for i := 1; i < len(envelopes); i++ {
		maxLis := math.MinInt
		for j := 0; j < i; j++ {
			if (envelopes[j][0] < envelopes[i][0]) && (envelopes[j][1] < envelopes[i][1]) {
				maxLis = max(maxLis, lis[j])
			}
		}

		if maxLis != math.MinInt {
			lis[i] += maxLis
		}
		ans = max(ans, lis[i])
	}

	return ans
}

func MaxEnvelopesv2(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	if len(envelopes) == 1 {
		return 1
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	var temp []int
	temp = append(temp, envelopes[0][1])

	for j := 1; j < len(envelopes); j++ {
		if envelopes[j][1] > temp[len(temp)-1] {
			temp = append(temp, envelopes[j][1])
		} else {
			index := findBS(temp, envelopes[j][1])
			temp[index] = envelopes[j][1]
		}
	}

	return len(temp)
}

func findBS(nums []int, val int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] == val {
			return mid
		} else if nums[mid] > val {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
