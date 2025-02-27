package leetcode

func TargetSumSubset(nums []int, target int) bool {
	n := len(nums)
	rows, cols := n+1, target+1
	dp := make([][]bool, rows)
	for i := range dp {
		dp[i] = make([]bool, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = false
			} else if j == 0 {
				dp[i][j] = true
			} else {
				if dp[i-1][j] {
					// if player i don't play to score j runs
					dp[i][j] = true
				} else {
					// if player i play to score runs = nums[i-1
					val := nums[i-1]
					if j >= val {
						if dp[i-1][j-val] {
							dp[i][j] = true
						}
					}
				}
			}
		}
	}

	return dp[n][target]
}
