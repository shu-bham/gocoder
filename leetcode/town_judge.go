package leetcode

// Easy: https://leetcode.com/problems/find-the-town-judge
func FindJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}

	degree := make([]int, n+1)
	for _, trust := range trust {
		from := trust[0]
		to := trust[1]
		degree[from]--
		degree[to]++
	}

	for i, deg := range degree {
		if deg == n-1 {
			return i
		}
	}

	return -1
}
