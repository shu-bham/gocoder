package leetcode

// Medium: https://leetcode.com/problems/minimum-index-of-a-valid-split

func MinimumIndexOfValidSplit(nums []int) int {
	n := len(nums)
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	dominant := -1
	for k, v := range freqMap {
		if v > n/2 {
			dominant = k
			break
		}
	}

	if dominant == -1 {
		return -1
	}

	leftCount, rightCount := 0, freqMap[dominant]
	for i, num := range nums {
		if num == dominant {
			leftCount++
			rightCount--
			if leftCount > (i+1)/2 && rightCount > (n-i-1)/2 {
				return i
			}
		}
	}

	return -1

}

// Boyer-Moore Majority Voting Algorithm
func MinimumIndexOfValidSplitV2(nums []int) int {
	n := len(nums)
	candidate := 0
	count := 0

	// Phase 1: Find a potential candidate
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	// Phase 2: Verify if the candidate is the majority element
	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}

	// Check if the candidate is actually the majority element
	if count <= n/2 {
		return -1
	}

	leftCount, rightCount := 0, count
	for i, num := range nums {
		if num == candidate {
			leftCount++
			rightCount--
			if leftCount > (i+1)/2 && rightCount > (n-i-1)/2 {
				return i
			}
		}
	}

	return -1

}
