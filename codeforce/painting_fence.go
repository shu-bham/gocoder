package codeforce

import (
	"fmt"
	"math"
)

// Painting Fence: https://codeforces.com/problemset/problem/448/C
func MinNoOfStrokes(heights []int) int {
	return helper(heights, 0, len(heights)-1)
}

func helper(heights []int, left int, right int) int {
	var returnVal int
	defer func() {
		fmt.Println(heights, left, right, returnVal)
	}()

	if right < left {
		returnVal = 0
		return returnVal
	}
	if left == right {
		if heights[left] == 0 {
			returnVal = 0
			return 0
		}
		returnVal = 1
		return 1
	}

	ans := 0
	minHeight := math.MaxInt
	for i := left; i <= right; i++ {
		minHeight = min(minHeight, heights[i])
		if heights[i] != 0 {
			ans++
		}
	}

	zeroPos := []int{}
	for i := left; i <= right; i++ {

		heights[i] -= minHeight
		if heights[i] == 0 {
			zeroPos = append(zeroPos, i)
		}
	}

	tempAns := 0
	start := left
	for i := range zeroPos {
		tempAns += helper(heights, start, zeroPos[i]-1)
		start = zeroPos[i] + 1
	}
	tempAns += helper(heights, start, right)

	returnVal = min(ans, minHeight+tempAns)
	return returnVal
}
