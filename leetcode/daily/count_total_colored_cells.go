package leetcode

// Medium: https://leetcode.com/problems/count-total-number-of-colored-cells
func ColoredCells(n int) int64 {
	if n == 1 {
		return 1
	}

	var prev, ans int64
	prev = 1
	for i := 2; i <= n; i++ {
		ans = prev + 4 + int64(4*(i-2))
		prev = ans
	}

	return ans

}

// optimised: there is a pattern forming for each n
func ColoredCellsV2(n int) int64 {
	if n == 1 {
		return 1
	}

	t := int(n)

	return int64(1 + 4*(t*(t-1)/2))

}
