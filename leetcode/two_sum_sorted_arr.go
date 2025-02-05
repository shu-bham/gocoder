package leetcode

func TwoSumSortedInput(numbers []int, target int) []int {

	lo, hi := 0, len(numbers)-1

	for lo < hi {
		s := numbers[lo] + numbers[hi]
		if s == target {
			return []int{lo + 1, hi + 1}
		} else if s > target {
			hi--
		} else {
			lo++
		}
	}

	return []int{-1, -1}

}
