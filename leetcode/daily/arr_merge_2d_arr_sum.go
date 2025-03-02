package leetcode

// Easy: https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/description/
func MergeArraysSum(nums1 [][]int, nums2 [][]int) [][]int {
	mm := make([]int, 1001)
	process(mm, nums1)
	process(mm, nums2)
	return convert(mm)
}

func convert(mm []int) [][]int {
	res := make([][]int, 0)
	for i, val := range mm {
		if val != 0 {
			res = append(res, []int{i, val})
		}
	}
	return res
}

func process(mm []int, arr [][]int) {
	for i := range arr {
		key := arr[i][0]
		val := arr[i][1]
		mm[key] += val
	}
}
