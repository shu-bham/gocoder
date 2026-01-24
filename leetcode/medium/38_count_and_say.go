package medium

import (
	"strconv"
	"strings"
)

/*
*
38. Count and Say
https://leetcode.com/problems/count-and-say

The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

countAndSay(1) = "1"
countAndSay(n) is the run-length encoding of countAndSay(n - 1).
Run-length encoding (RLE) is a string compression method that works by replacing consecutive identical characters (repeated 2 or more times) with the concatenation of the character and the number marking the count of the characters (length of the run). For example, to compress the string "3322251" we replace "33" with "23", replace "222" with "32", replace "5" with "15" and replace "1" with "11". Thus the compressed string becomes "23321511".

Given a positive integer n, return the nth element of the count-and-say sequence.

'Adobe', 'Amazon', 'Apple', 'Bloomberg', 'Epic-systems', 'Facebook', 'Google', 'Microsoft', 'Oracle
*/

func CountAndSay(n int) string {
	dp := make(map[int]string)
	dp[1] = "1"
	return f(n, dp)
}

func f(n int, dp map[int]string) string {
	s, ok := dp[n]
	if ok {
		return s
	}

	say := f(n-1, dp)
	dp[n] = rle(say)
	return dp[n]
}

func rle(str string) string {
	if str == "" {
		return ""
	}
	var ans strings.Builder
	runes := []rune(str)
	currCount := 1
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			currCount++
		} else {
			ans.WriteString(strconv.Itoa(currCount))
			ans.WriteRune(runes[i-1])
			currCount = 1
		}

	}
	ans.WriteString(strconv.Itoa(currCount))
	ans.WriteRune(runes[len(runes)-1])
	return ans.String()
}

func CountAndSayIter(n int) string {
	ans := "1"
	if n == 1 {
		return ans
	}

	for i := 2; i <= n; i++ {
		ans = rle(ans)
	}

	return ans
}
