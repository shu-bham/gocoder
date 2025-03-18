package leetcode

// Easy: https://leetcode.com/problems/longest-nice-subarray
func LongestNiceSubarray(nums []int) int {
	res := 1
	for i := 0; i < len(nums); i++ {
		addSum, orSum := nums[i], nums[i]
		for j := i + 1; j < len(nums); j++ {
			addSum += nums[j]
			orSum |= nums[j]

			if orSum == addSum {
				res = max(res, j-i+1)
				continue
			} else {
				break
			}
		}

	}
	return res
}

func LongestNiceSubarrayV2(nums []int) int {
	left, res := 0, 1
	addSum, orSum := 0, 0

	for right := 0; right < len(nums); right++ {
		addSum += nums[right]
		orSum |= nums[right]

		// If OR sum and add sum are not equal, move left pointer
		for left <= right && addSum != orSum {
			addSum -= nums[left]

			// Recalculate OR sum by recomputing from left to right
			orSum = 0
			for i := left + 1; i <= right; i++ {
				orSum |= nums[i]
			}

			left++
		}

		res = max(res, right-left+1)
	}

	return res
}
