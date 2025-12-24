package google

// brute force : n^3
func LengthOfLongestSubstring(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			freqMap := make(map[uint8]int)
			invalid := false
			for k := i; k <= j; k++ {
				ch := s[k]
				_, ok := freqMap[ch]
				if !ok {
					freqMap[ch]++
				} else {
					invalid = true
					break
				}
			}

			if !invalid {
				ans = max(ans, len(freqMap))
			}
		}
	}
	return ans
}

// sliding window
func LengthOfLongestSubstringV2(s string) int {
	ans := 0
	i := 0
	charPos := make(map[uint8]int)
	for j := 0; j < len(s); j++ {
		ch := s[j]
		prevIndex, ok := charPos[ch]
		if ok && prevIndex >= i {
			i = prevIndex + 1
		}
		charPos[ch] = j
		ans = max(ans, j-i+1)

	}

	return ans
}
