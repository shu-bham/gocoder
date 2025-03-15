package leetcode

// Medium: https://leetcode.com/problems/house-robber/
func Rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	if n == 1 {
		return nums[0]
	}

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		// if robs i'th house
		profitIfRobs := nums[i] + dp[i-2]
		// if doesn't robs i'th house
		profitIfNotRobs := dp[i-1]
		dp[i] = max(profitIfRobs, profitIfNotRobs)
	}
	return dp[n-1]
}
