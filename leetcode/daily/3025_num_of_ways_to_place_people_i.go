package leetcode

import (
	"math"
	"sort"
)

/**
https://leetcode.com/problems/find-the-number-of-ways-to-place-people-i/description/?envType=daily-question&envId=2025-09-02

You are given a 2D array points of size n x 2 representing integer coordinates of some points on a 2D plane, where points[i] = [xi, yi].

Count the number of pairs of points (A, B), where

A is on the upper left side of B, and
there are no other points in the rectangle (or line) they make (including the border).
Return the count.

*/

func NumberOfPairs(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		first := points[i]
		second := points[j]
		if first[0] != second[0] {
			return first[0] <= second[0]
		} else {
			return first[1] > second[1]
		}
	})

	ans := 0

	for i := 0; i < len(points); i++ {
		left := points[i]
		upperY := left[1]
		lowerYLimit := math.MinInt
		for j := i + 1; j < len(points); j++ {
			right := points[j]
			currY := right[1]

			if currY <= upperY && currY > lowerYLimit {
				ans++
				lowerYLimit = currY
				if currY == upperY {
					// since next all points will cause this point to be a part of the rectangle
					break
				}
			}

		}
	}

	return ans
}
