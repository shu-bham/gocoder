package leetcode

// Easy: https://leetcode.com/problems/find-missing-and-repeated-values
func FindMissingAndRepeatedValues(grid [][]int) []int {
	r, c := len(grid), len(grid[0])

	dp := make([]int, 2501)
	for i := range dp {
		dp[i] = -1
	}
	var repeat, missing int
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			num := grid[i][j]
			if dp[num] != -1 {
				repeat = num
			} else {
				dp[num] = 1
			}
		}
	}

	for i := 1; i < len(dp); i++ {
		if dp[i] == -1 {
			missing = i
			break
		}
	}

	return []int{repeat, missing}
}

func FindMissingAndRepeatedValuesV2(grid [][]int) []int {
	n := len(grid) * len(grid)

	expected_sum := n * (n + 1) / 2
	expected_sum_squares := n * (n + 1) * (2*n + 1) / 6

	actual_sum := 0
	actual_sum_squares := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			actual_sum += grid[i][j]
			actual_sum_squares += (grid[i][j] * grid[i][j])
		}
	}

	// a-b
	diff_sum := actual_sum - expected_sum
	// a^2 - b^2
	diff_sum_squares := actual_sum_squares - expected_sum_squares

	// a+b
	sum_a_b := diff_sum_squares / diff_sum

	a := (sum_a_b + diff_sum) / 2
	b := (sum_a_b - diff_sum) / 2
	return []int{a, b}

}
