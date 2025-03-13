package leetcode

// Medium: https://leetcode.com/problems/zero-array-transformation-ii
func MinZeroArray(nums []int, queries [][]int) int {
	left, right := 0, len(queries)

	if !zeroSumArrayFormed(nums, queries, right) {
		return -1
	}
	for left <= right {
		mid := (left + right) / 2
		if zeroSumArrayFormed(nums, queries, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func zeroSumArrayFormed(nums []int, queries [][]int, k int) bool {
	n := len(nums)
	differenceArr := make([]int, n+1)

	for i := 0; i < k; i++ {
		left := queries[i][0]
		right := queries[i][1]
		val := queries[i][2]

		differenceArr[left] += val
		differenceArr[right+1] -= val
	}

	sum := 0
	// calc prefix sum from differenceArr
	for i := 0; i < n; i++ {
		sum += differenceArr[i]
		if sum < nums[i] {
			return false
		}
	}
	return true

}
