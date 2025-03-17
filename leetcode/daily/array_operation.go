package leetcode

// Easy: https://leetcode.com/problems/apply-operations-to-an-array
// Bruteforce
func ApplyOperations(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = 2 * nums[i]
			nums[i+1] = 0
			i++
		}
	}

	i := 0
	for _, num := range nums {
		if num != 0 {
			res[i] = num
			i++
		}
	}

	return res

}

// Space optimised: in-place arrangement
func ApplyOperationsV2(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = 2 * nums[i]
			nums[i+1] = 0
			i++
		}
	}

	nonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex] = nums[i]
			nonZeroIndex++
		}
	}

	for nonZeroIndex < len(nums) {
		nums[nonZeroIndex] = 0
		nonZeroIndex++
	}

	return nums
}
