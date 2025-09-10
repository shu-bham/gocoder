package hard

import "math"

/**
1931. Painting a Grid With Three Different Colors
https://leetcode.com/problems/painting-a-grid-with-three-different-colors

You are given two integers m and n. Consider an m x n grid where each cell is initially white. You can paint each cell red, green, or blue. All cells must be painted.

Return the number of ways to color the grid with no two adjacent cells having the same color. Since the answer can be very large, return it modulo 109 + 7.
*/

const mod = 1000000007

func ColorTheGrid(m int, n int) int {
	// Hash mapping stores all valid coloration schemes for a single row that
	// meet the requirements
	valid := make(map[int][]int)

	// Enumerate masks that meet the requirements within the range [0, 3^m)
	maskEnd := int(math.Pow(3, float64(m)))
	for mask := 0; mask < maskEnd; mask++ {
		color := make([]int, m)
		mm := mask
		for i := 0; i < m; i++ {
			color[i] = mm % 3
			mm /= 3
		}
		check := true
		for i := 0; i < m-1; i++ {
			if color[i] == color[i+1] {
				check = false
				break
			}
		}
		if check {
			valid[mask] = color
		}
	}

	// Preprocess all (mask1, mask2) binary tuples, satisfying mask1 and mask2
	// When adjacent rows, the colors of the two cells in the same column are
	// different
	adjacent := make(map[int][]int)
	for mask1 := range valid {
		for mask2 := range valid {
			check := true
			for i := 0; i < m; i++ {
				if valid[mask1][i] == valid[mask2][i] {
					check = false
					break
				}
			}
			if check {
				adjacent[mask1] = append(adjacent[mask1], mask2)
			}
		}
	}

	f := make(map[int]int)
	// for 1st col
	for mask := range valid {
		f[mask] = 1
	}

	for i := 1; i < n; i++ {
		g := make(map[int]int)
		// mask2 is the color combination for the current row (ith col)
		for mask2 := range valid {
			// mask1 is the color combination of the previous row (i-1th col)
			for _, mask1 := range adjacent[mask2] {
				// no of ways to make mask2 color combinations is equal to all ways of making mask1 in prev column
				g[mask2] = (g[mask2] + f[mask1]) % mod
			}
		}
		// previous = current for next iteration
		f = g
	}

	ans := 0
	// sum of all values in f is the total color combinations
	for _, num := range f {
		ans = (ans + num) % mod
	}
	return ans
}
