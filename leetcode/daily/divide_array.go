package leetcode

// Easy: https://leetcode.com/problems/divide-array-into-equal-pairs
func DivideArray(nums []int) bool {
	if len(nums)%2 != 0 {
		return false
	}

	freq := make(map[uint]uint)
	for _, num := range nums {
		freq[uint(num)]++
	}

	for _, v := range freq {
		if v%2 != 0 {
			return false
		}
	}
	return true
}
