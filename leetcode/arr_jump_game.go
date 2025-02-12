package leetcode

func CanJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)

	for i := n - 1; i >= 0; i-- {
		//fmt.Println(nums[i], dp)
		jump := nums[i]
		if i == n-1 {
			dp[i] = true
		} else {
			for j := i + 1; j <= i+jump && j < n; j++ {
				if j == n-1 || dp[j] {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[0]
}

func CanJumpV2(nums []int) bool {
	maxJump := 0
	for _, val := range nums {
		if maxJump < 0 {
			return false
		} else if val > maxJump {
			maxJump = val
		}
		//fmt.Println(val, maxJump)
		maxJump--
	}
	return true

}
