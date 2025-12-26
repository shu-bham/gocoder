package google

func LongestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	var start, end int

	for gap := 0; gap < len(s); gap++ {
		i := 0
		j := gap
		for j < len(s) {
			if gap == 0 {
				dp[i][j] = true
			} else if gap == 1 {
				if s[i] == s[j] {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			} else {
				if s[i] == s[j] && dp[i+1][j-1] == true {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			}
			if dp[i][j] {
				start = i
				end = j
			}

			i++
			j++
		}
	}
	return s[start : end+1]
}
