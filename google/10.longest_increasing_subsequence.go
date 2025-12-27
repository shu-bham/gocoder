package google

import (
	"math"
)

func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lis := make([]int, len(nums))
	lis[0] = 1

	for i := 1; i < len(nums); i++ {
		maxLIS := math.MinInt
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				maxLIS = max(maxLIS, lis[j])
			}
		}
		if maxLIS == math.MinInt {
			lis[i] = 1
		} else {
			lis[i] = 1 + maxLIS
		}
	}

	ans := math.MinInt
	for _, l := range lis {
		ans = max(ans, l)
	}

	return ans
}

func LengthOfLISv2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var temp []int
	temp = append(temp, nums[0])

	for i := 1; i < len(nums); i++ {
		if nums[i] > temp[len(temp)-1] {
			temp = append(temp, nums[i])
		} else {
			index := find(temp, nums[i])
			temp[index] = nums[i]
		}
	}

	return len(temp)
}

func find(temp []int, num int) int {
	left, right := 0, len(temp)-1

	for left < right {
		mid := (left + right) / 2
		if temp[mid] == num {
			return mid
		} else if temp[mid] > num {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
