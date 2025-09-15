package hard

/**
https://leetcode.com/problems/unique-paths-iii
You are given an m x n integer array grid where grid[i][j] could be:
1 representing the starting square. There is exactly one starting square.
2 representing the ending square. There is exactly one ending square.
0 representing empty squares we can walk over.
-1 representing obstacles that we cannot walk over.
Return the number of 4-directional walks from the starting square to the ending square, that walk over every non-obstacle square exactly once.

'Amazon', 'Google', 'Limebike
*/

func UniquePathsIII(grid [][]int) int {

	var empty = 1 // count of empty cells (0-cells)
	var res int

	var sx, sy int // pos of start point
	r, c := len(grid), len(grid[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 1 {
				sx, sy = i, j
			} else if grid[i][j] == 0 {
				empty++
			}
		}
	}

	// DFS + backtracking
	dfs(grid, sx, sy, &res, &empty)
	return res

}

func dfs(grid [][]int, x int, y int, res *int, empty *int) {
	// boundary conditions + skip obstacles
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] < 0 {
		return
	}

	if grid[x][y] == 2 {
		if *empty == 0 {
			// when all cells are covered
			*res++
		}
		return
	}

	grid[x][y] = -2
	*empty--
	dfs(grid, x, y+1, res, empty)
	dfs(grid, x+1, y, res, empty)
	dfs(grid, x-1, y, res, empty)
	dfs(grid, x, y-1, res, empty)
	grid[x][y] = 0
	*empty++
}
