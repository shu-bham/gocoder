package google

func MaxArea(height []int) int {
	ans := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			water := (j - i) * min(height[i], height[j])
			ans = max(ans, water)
		}
	}
	return ans
}

func MaxAreaV2(height []int) int {
	ans := 0
	i, j := 0, len(height)-1
	for i < j {
		water := (j - i) * min(height[i], height[j])
		ans = max(ans, water)
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}
