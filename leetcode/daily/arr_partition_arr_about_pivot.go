package leetcode

// Medium : https://leetcode.com/problems/partition-array-according-to-given-pivot
func PivotArray(nums []int, pivot int) []int {
	var small, large []int
	for _, num := range nums {
		if num < pivot {
			small = append(small, num)
		} else if num > pivot {
			large = append(large, num)
		}
	}

	pivotCount := len(nums) - len(small) - len(large)
	i := 0
	for _, num := range small {
		nums[i] = num
		i++
	}
	for j := 0; j < pivotCount; j++ {
		nums[i] = pivot
		i++
	}
	for _, num := range large {
		nums[i] = num
		i++
	}

	return nums
}

// space optimised
func PivotArrayV2(nums []int, pivot int) []int {
	res := make([]int, len(nums))
	lo, hi := 0, len(nums)-1

	for i := 0; i < len(nums); i++ {
		res[i] = pivot
		if nums[i] < pivot {
			res[lo] = nums[i]
			lo++
		}
	}

	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] > pivot {
			res[hi] = nums[j]
			hi--
		}
	}

	return res
}
