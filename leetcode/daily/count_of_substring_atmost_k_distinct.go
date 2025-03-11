package leetcode

func CountOfSubstringsAtmostK(word string, k int) int {
	n := len(word)
	left := 0
	res := 0
	mm := make(map[byte]int) // Use byte instead of uint8 (same thing but more idiomatic)

	for right := 0; right < n; right++ {
		mm[word[right]]++ // Directly increment without checking existence

		// Shrink window if we exceed k distinct characters
		for len(mm) > k {
			mm[word[left]]--
			if mm[word[left]] == 0 {
				delete(mm, word[left])
			}
			left++
		}

		// Count valid substrings ending at `right`
		res += right - left + 1
	}

	return res
}
