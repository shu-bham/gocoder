package google

import "math"

func MinNumberOperations(target []int) int {
	n := len(target)
	return countways(target, 0, n-1)
}

func countways(target []int, i int, j int) int {
	start := i
	end := j
	if i < 0 || i >= len(target) || j < 0 || j >= len(target) {
		return 0
	}
	if i == j {
		return target[i]
	}

	var zeroPos []int
	minEle := math.MaxInt
	ans := 0
	for index, elem := range target {
		if elem == 0 {
			zeroPos = append(zeroPos, index)
		}
		minEle = min(minEle, elem)
	}

	var points [][]int
	for _, index := range zeroPos {
		points = append(points, []int{start, index - 1})
		start = index + 1
	}
	if start != i {
		points = append(points, []int{start, end})
	}

	for _, ints := range points {
		ans += countways(target, ints[0], ints[1])
	}

	return minEle + countways(target, i, j)

}
