package google

import (
	"sort"
)

// 16ms, 4MB
func NextBeautifulNumberv2(n int) int {
	i := n + 1
	for {
		if isBeautifulNumber(i) {
			return i
		}
		i++
	}
}

func isBeautifulNumber(n int) bool {
	freq := make([]int, 9)
	for n > 0 {
		if n%10 == 0 {
			return false
		}
		freq[n%10-1]++
		n /= 10
	}
	for i := 0; i < 9; i++ {
		if freq[i] != 0 && freq[i] != i+1 {
			return false
		}
	}
	return true
}

// 500ms, 5MB
func NextBeautifulNumber(n int) int {
	freqMap := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var list []int
	generateAllNo(0, freqMap, &list)
	sort.Ints(list)
	for _, num := range list {
		if num > n {
			return num
		}
	}
	return -1
}

func generateAllNo(num int, freqMap []int, list *[]int) {
	if num > 0 && isBeautiful(freqMap) {
		*list = append(*list, num)
	}
	if num > 1224444 {
		return
	}

	for d := 1; d <= 7; d++ {
		if freqMap[d] < d {
			newNum := num*10 + d
			freqMap[d]++
			generateAllNo(newNum, freqMap, list)
			freqMap[d]--
		}
	}
}

func isBeautiful(freqMap []int) bool {
	for d := 1; d <= 7; d++ {
		freq := freqMap[d]
		if freq != 0 && freq != d {
			return false
		}
	}
	return true
}
