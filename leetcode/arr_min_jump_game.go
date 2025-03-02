package leetcode

func MinJump(nums []int) int {
	const MAX = 10001
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = MAX
	}
	dp[n-1] = 0

	for i := n - 2; i >= 0; i-- {
		jump := nums[i]
		for j := i + 1; j < n && j <= i+jump; j++ {
			dp[i] = min(dp[i], 1+dp[j])
		}

	}

	if dp[0] == MAX {
		return 0
	}

	return dp[0]
}
