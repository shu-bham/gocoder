package google

import (
	"golang.org/x/exp/slices"
)

func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	var ans [][]int
	var currentAns []int
	var backtrack func(index int, sumSoFar int)

	backtrack = func(index int, sum int) {
		if index == len(candidates) {
			if sum == 0 {
				newAns := make([]int, len(currentAns))
				copy(newAns, currentAns)
				ans = append(ans, newAns)
			}
			return
		}
		if sum-candidates[index] >= 0 {
			currentAns = append(currentAns, candidates[index])
			// take
			backtrack(index, sum-candidates[index])
			currentAns = currentAns[:len(currentAns)-1]
		}
		// don't take
		backtrack(index+1, sum)
	}

	backtrack(0, target)
	return ans
}
