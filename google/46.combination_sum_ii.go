package google

import (
	"slices"
)

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	var ans [][]int
	var currAns []int
	var backtrack func(index int, target int)
	backtrack = func(index int, target int) {
		if target == 0 {
			temp := make([]int, len(currAns))
			copy(temp, currAns)
			ans = append(ans, temp)
			return
		}

		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i] > target {
				break
			}
			currAns = append(currAns, candidates[i])
			backtrack(i+1, target-candidates[i])
			currAns = currAns[:len(currAns)-1]
		}
	}
	backtrack(0, target)
	//fmt.Println(ans)
	return ans
}
