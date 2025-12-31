package google

func CanSeePersonsCount(heights []int) []int {
	type ele struct {
		ht, index int
	}
	var stack []ele
	res := make([]int, len(heights))
	for i, height := range heights {
		for len(stack) > 0 && stack[len(stack)-1].ht < height {
			top := stack[len(stack)-1]
			res[top.index]++
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			top := stack[len(stack)-1]
			res[top.index]++
		}

		stack = append(stack, ele{height, i})
	}

	return res
}
