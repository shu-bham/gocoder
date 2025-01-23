package majority_element

import "sort"

// Majority Element
func MajorityElement(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	n := len(nums)
	return nums[n/2]
}
