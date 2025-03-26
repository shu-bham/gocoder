package leetcode

import (
	"log"
	"math"
)

type obj struct {
	val int
	ok  bool
}

// Hard: https://leetcode.com/problems/word-ladder/
// todo: not optmised, getting TLE
func LadderLength(beginWord string, endWord string, wordList []string) int {

	visited := make(map[string]bool)
	dp := make(map[string]obj)
	len, ok := getLen(beginWord, endWord, wordList, visited, dp)
	if ok {
		return len
	}

	return 0
}

func getLen(begin string, end string, wordList []string, visited map[string]bool, dp map[string]obj) (int, bool) {
	visited[begin] = true
	defer func() {
		visited[begin] = false
	}()

	val, ok := dp[begin+end]
	if ok {
		log.Default().Println(begin, end, "DP_LOOKUP", val)
		return val.val, val.ok
	}
	if begin == end {
		log.Default().Println(begin, end, "reached end", 1)
		dp[begin+end] = obj{1, true}
		return 1, true
	}

	next := getNextWord(begin, wordList, visited)
	minLen := math.MaxInt
	for _, s := range next {
		l, ok := getLen(s, end, wordList, visited, dp)
		if ok {
			minLen = min(minLen, l)
		}
	}

	// if minLen not found then unreachable to end
	if minLen == math.MaxInt {
		log.Default().Println(begin, end, "unreachable", -1)
		return -1, false
	}
	log.Default().Println(begin, end, "getLen:", 1+minLen)
	dp[begin+end] = obj{1 + minLen, true}
	return 1 + minLen, true

}

func getNextWord(curr string, list []string, visited map[string]bool) []string {
	numOfDiffChar := func(a, b string) int {
		ctr := 0
		n := len(a)
		for i := 0; i < n; i++ {
			if a[i] != b[i] {
				ctr++
			}
		}
		return ctr
	}

	var res []string
	for _, word := range list {
		if visited[word] {
			continue
		}
		if numOfDiffChar(curr, word) == 1 {
			res = append(res, word)
		}
	}
	//log.Default().Println("[nextWord]", curr, res)
	return res
}
