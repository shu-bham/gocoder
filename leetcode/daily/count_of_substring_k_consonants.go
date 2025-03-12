package leetcode

// Medium: https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-ii
func CountOfSubstrings(word string, k int) int64 {
	start, end := 0, 0
	consonantCount := 0
	vowMap := make(map[uint8]int)
	res := int64(00)
	n := len(word)
	nextConsonant := make([]int, n)

	nextConsonantIndex := n
	for i := n - 1; i >= 0; i-- {
		nextConsonant[i] = nextConsonantIndex
		if !isVowel(word[i]) {
			nextConsonantIndex = i
		}
	}

	for end < n {
		newChar := word[end]
		if isVowel(newChar) {
			vowMap[newChar]++
		} else {
			consonantCount++
		}

		// shrink
		for consonantCount > k {
			startCh := word[start]
			if isVowel(startCh) {
				vowMap[startCh]--
				if vowMap[startCh] == 0 {
					delete(vowMap, startCh)
				}
			} else {
				consonantCount--
			}
			start++
		}

		for consonantCount == k && len(vowMap) == 5 {
			res += int64(nextConsonant[end] - end)
			startCh := word[start]
			if isVowel(startCh) {
				vowMap[startCh]--
				if vowMap[startCh] == 0 {
					delete(vowMap, startCh)
				}
			} else {
				consonantCount--
			}
			start++
		}

		end++
	}

	return res
}

func CountOfSubstringsV2(word string, k int) int64 {
	return atLeast(word, k) - atLeast(word, k+1)
}

func atLeast(word string, k int) int64 {
	start, end := 0, 0
	vowMap := make(map[uint8]int)
	consonantCount := 0
	res := int64(0)
	n := len(word)
	for end < n {
		newChar := word[end]
		if isVowel(newChar) {
			vowMap[newChar]++
		} else {
			consonantCount++
		}

		for len(vowMap) == 5 && consonantCount >= k {
			res += int64(n - end)
			startCh := word[start]
			if isVowel(startCh) {
				vowMap[startCh]--
				if vowMap[startCh] == 0 {
					delete(vowMap, startCh)
				}
			} else {
				consonantCount--
			}
			start++
		}
		end++
	}
	return res
}

func isVowel(ch uint8) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}
