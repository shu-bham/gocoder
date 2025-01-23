package leetcode

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		v := target - num
		j, ok := m[v]
		if ok {
			if i < j {
				return []int{i, j}
			}
			return []int{j, i}
		}
		m[num] = i
	}

	return nil

}
