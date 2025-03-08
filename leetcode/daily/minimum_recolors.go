package leetcode

// Easy: https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks
func MinimumRecolors(blocks string, k int) int {
	black_count, white_count := 0, 0
	recolors := 0
	for i := 0; i < k; i++ {
		ch := blocks[i]
		if ch == 'B' {
			black_count++
		} else {
			white_count++
		}
	}

	recolors = white_count

	for i := k; i < len(blocks); i++ {
		curr := blocks[i]
		prev := blocks[i-k]

		if curr == 'B' {
			black_count++
		} else {
			white_count++
		}
		if prev == 'B' {
			black_count--
		} else {
			white_count--
		}

		recolors = min(recolors, white_count)
	}

	return recolors
}
