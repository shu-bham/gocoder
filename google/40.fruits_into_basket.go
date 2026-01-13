package google

func TotalFruit(fruits []int) int {
	n := len(fruits)
	// find the max window of fruits having only 2 kinds of trees
	i := 0
	j := 0
	ans := 0
	freqMap := make(map[int]int)
	for j < n {
		curr := fruits[j]
		freqMap[curr]++
		for len(freqMap) > 2 {
			// shorten
			start := fruits[i]
			freqMap[start]--
			i++
			if freqMap[start] == 0 {
				delete(freqMap, start)
			}
		}
		ans = max(ans, j-i+1)
		j++
	}

	return ans
}
