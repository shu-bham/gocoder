package leetcode

// Medium: https://leetcode.com/problems/alternating-groups-ii
// sliding window
func NumberOfAlternatingGroups(colors []int, k int) int {
	l, r := 0, 0
	ctr := 0

	n := len(colors)
	for l < n {
		r++
		if colors[r%n] == colors[(r-1)%n] {
			l = r
		} else if r-l+1 == k {
			ctr++
			l++
		}
	}
	return ctr
}
