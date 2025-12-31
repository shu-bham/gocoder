package google

// monotonic stack
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	nge := make([]int, len(nums2))

	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			nge[i] = stack[len(stack)-1]
		} else {
			nge[i] = -1
		}

		stack = append(stack, nums2[i])
	}

	posMap := make(map[int]int)
	for j, num := range nums2 {
		posMap[num] = j
	}

	ans := make([]int, len(nums1))
	for i, num := range nums1 {
		index := posMap[num]
		nbr := nge[index]
		ans[i] = nbr
	}

	return ans
}
