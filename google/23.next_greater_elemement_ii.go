package google

import (
	"math"
)

func NextGreaterElements(nums []int) []int {
	maxPos := -1
	maxSoFar := math.MinInt

	N := len(nums)
	nge := make([]int, N)

	for i := range nums {
		if nums[i] > maxSoFar {
			maxSoFar = nums[i]
			maxPos = i
		}
	}

	var stack []int
	for i := 0; i < N; i++ {
		currIndex := (maxPos - i + N) % N

		for len(stack) > 0 && stack[len(stack)-1] <= nums[currIndex] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			nge[currIndex] = stack[len(stack)-1]
		} else {
			nge[currIndex] = -1
		}

		stack = append(stack, nums[currIndex])
	}
	return nge
}
