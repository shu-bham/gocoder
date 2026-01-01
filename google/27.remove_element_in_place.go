package google

func RemoveElement(nums []int, val int) int {
	pos := 0
	for _, n := range nums {
		if n == val {
			continue
		} else {
			nums[pos] = n
			pos++
		}
	}
	return pos
}
