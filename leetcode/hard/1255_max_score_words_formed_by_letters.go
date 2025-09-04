package hard

// MaxScoreWords
// Let W be the length of words
// L be the maximum length of any word in words,
// and A be the size of the alphabet (in this case, A=26).
// Time complexity: O(2^W * (WL+A)).
// Space complexity: O(A)
func MaxScoreWords(words []string, letters []byte, score []int) int {
	n := len(words)
	ans := 0
	letterMap := createLetterMap(letters)
	// check all permutations of words
	for mask := 0; mask < (1 << n); mask++ {
		subsetLetterMap := make(map[byte]int)
		for i := 0; i < n; i++ {
			// check if ith bit of mask is set
			if mask&(1<<i) > 0 {
				addToMap(subsetLetterMap, words[i])
			}
		}
		if isValid(subsetLetterMap, letterMap) {
			ans = max(ans, calcScore(subsetLetterMap, score))
		}

	}
	return ans
}

// check if letterMap contains all letters and frequency as that in subsetLetterMap
func isValid(subsetLetterMap map[byte]int, letterMap map[byte]int) bool {
	for ch, freq := range subsetLetterMap {
		count, ok := letterMap[ch]
		if !ok {
			return false
		}
		if freq > count {
			return false
		}
	}
	return true
}

func calcScore(letterMap map[byte]int, score []int) int {
	sum := 0
	for ch, freq := range letterMap {
		sum += score[ch-'a'] * freq
	}
	return sum
}

func createLetterMap(letters []byte) map[byte]int {
	res := make(map[byte]int)
	for _, letter := range letters {
		res[letter]++
	}
	return res
}

func addToMap(letterMap map[byte]int, word string) {
	for _, letter := range word {
		letterMap[byte(letter)]++
	}
}
