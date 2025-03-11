package leetcode

// Medium: https://leetcode.com/problems/number-of-substrings-containing-all-three-characters
// bruteforce
func CountOfSubstringsExactK(word string, k int) int {
	n := len(word)
	res := int(0)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			mm := make(map[uint8]int)
			for p := i; p <= j; p++ {
				curr := word[p]
				mm[curr] = 1
			}

			if len(mm) == k {
				res++
			}

		}
	}
	return res
}

// optimised: sliding window
func CountOfSubstringsExactKV2(word string, k int) int {
	left, right := 0, 0
	mm := make(map[uint8]int)
	res := 0
	n := len(word)
	for right < n {
		curr := word[right]
		v, ok := mm[curr]
		if ok {
			mm[curr] = v + 1
		} else {
			mm[curr] = 1
		}

		for len(mm) == k && left <= right {
			res += n - right
			val := mm[word[left]]
			if val == 1 {
				delete(mm, word[left])
			} else {
				mm[word[left]] = val - 1
			}
			left++
		}

		right++
	}
	return res
}
