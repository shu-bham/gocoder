package google

// for each i, find left max and right max, bruteforce
func Trap(height []int) int {
	ans := 0
	for i := 1; i < len(height)-1; i++ {
		leftMax := height[i]
		for j := 0; j < i; j++ {
			leftMax = max(leftMax, height[j])
		}

		rightMax := height[i]
		for j := i + 1; j < len(height); j++ {
			rightMax = max(rightMax, height[j])
		}

		ans += min(leftMax, rightMax) - height[i]
	}

	return ans
}

func TrapV2(height []int) int {
	ans := 0
	n := len(height)
	maxToLeft := make([]int, n)
	maxToRight := make([]int, n)

	maxToLeft[0] = height[0]
	for i := 1; i < n; i++ {
		maxToLeft[i] = max(maxToLeft[i-1], height[i])
	}

	maxToRight[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		maxToRight[i] = max(maxToRight[i+1], height[i])
	}

	for i := 0; i < n; i++ {
		ans += min(maxToLeft[i], maxToRight[i]) - height[i]
	}

	return ans
}
