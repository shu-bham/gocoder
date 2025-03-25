package leetcode

import (
	"sort"
)

// Medium: https://leetcode.com/problems/check-if-grid-can-be-cut-into-sections

type Point struct {
	x, y int
}
type Rect struct {
	bottomLeft, topRight Point
}

func CheckValidCuts(n int, rectangles [][]int) bool {
	rect := createRectangles(rectangles)
	var xpoints, ypoints [][]int

	for _, r := range rect {
		xpoints = append(xpoints, r.getXCoord())
		ypoints = append(ypoints, r.getYCoord())
	}

	xintervals := mergeOverlapIntervals(xpoints)
	if xintervals >= 3 {
		return true
	}

	yintervals := mergeOverlapIntervals(ypoints)
	if yintervals >= 3 {
		return true
	}

	return false
}

func createRectangles(rectangles [][]int) []Rect {
	var res []Rect
	for _, rect := range rectangles {
		res = append(res, Rect{
			bottomLeft: Point{rect[0], rect[1]},
			topRight:   Point{rect[2], rect[3]},
		})
	}
	return res
}

func (r Rect) getXCoord() []int {
	return []int{r.bottomLeft.x, r.topRight.x}
}

func (r Rect) getYCoord() []int {
	return []int{r.bottomLeft.y, r.topRight.y}
}

func mergeOverlapIntervals(intervals [][]int) int {

	if len(intervals) <= 1 {
		return len(intervals)
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	lastMergedIndex := 0

	for i := 1; i < len(intervals); i++ {
		merged := intervals[lastMergedIndex]
		curr := intervals[i]

		if curr[0] < merged[1] {
			merged[1] = max(merged[1], curr[1])
		} else {
			lastMergedIndex++
			intervals[lastMergedIndex] = intervals[i]
		}
	}

	return lastMergedIndex + 1
}

type pair struct{ l, r int }

// fast
func CheckValidCutsV2(n int, rectangles [][]int) bool {
	xp := make([]pair, n)
	yp := make([]pair, n)
	for i, r := range rectangles {
		xp[i] = pair{r[0], r[2]}
		yp[i] = pair{r[1], r[3]}
	}

	return check(xp) || check(yp)
}

func check(pairs []pair) bool {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].l < pairs[j].l
	})

	count := 0
	maxSoFar := 0

	for i := 0; i < len(pairs); i++ {
		curr := pairs[i]
		if curr.l >= maxSoFar {
			count++
		}
		maxSoFar = max(maxSoFar, curr.r)
	}

	return count >= 3

}
