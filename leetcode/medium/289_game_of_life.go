package medium

/**
https://leetcode.com/problems/game-of-life/
289. Game of Life

According to Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

The board is made up of an m x n grid of cells, where each cell has an initial state: live (represented by a 1) or dead (represented by a 0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

Any live cell with fewer than two live neighbors dies as if caused by under-population.
Any live cell with two or three live neighbors lives on to the next generation.
Any live cell with more than three live neighbors dies, as if by over-population.
Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
The next state of the board is determined by applying the above rules simultaneously to every cell in the current state of the m x n grid board. In this process, births and deaths occur simultaneously.

Given the current state of the board, update the board to reflect its next state.

Note that you do not need to return anything.

'Amazon', 'Bloomberg', 'Dropbox', 'Facebook', 'Goldman-sachs', 'Google', 'Microsoft', 'Opendoor', 'Oracle', 'Reddit', 'Snapchat', 'Two-sigma', 'Uber', 'Zillow'
*/

// 2 boards
func GameOfLifeV1(board [][]int) {
	m, n := len(board), len(board[0])
	nextBoard := make([][]int, m)
	for i := 0; i < m; i++ {
		nextBoard[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			curr := board[i][j]
			ctr := countNbr(board, i, j)
			nextBoard[i][j] = nextStateOfCell(curr, ctr)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = nextBoard[i][j]
		}
	}
}

// single board
func GameOfLifeOptimised(board [][]int) {
	m, n := len(board), len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			curr := board[i][j]
			ctr := countNbrNew(board, i, j)
			board[i][j] = nextStateOfCellNew(curr, ctr)
		}
	}

	// reset nums for live/ dead
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}

}

func countNbr(board [][]int, i int, j int) int {
	m, n := len(board), len(board[0])
	ctr := 0
	nbrPos := [][]int{{i - 1, j}, {i - 1, j + 1}, {i, j + 1}, {i + 1, j + 1}, {i + 1, j}, {i + 1, j - 1}, {i, j - 1}, {i - 1, j - 1}}
	for _, pos := range nbrPos {
		x := pos[0]
		y := pos[1]
		if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 1 {
			ctr++
		}
	}
	return ctr
}

func nextStateOfCell(curr int, liveNbr int) int {
	if curr == 1 {
		if liveNbr < 2 {
			return 0
		} else if liveNbr < 4 {
			return 1
		} else {
			return 0
		}
	} else {
		if liveNbr == 3 {
			return 1
		} else {
			return 0
		}
	}
}

func countNbrNew(board [][]int, i int, j int) int {
	m, n := len(board), len(board[0])
	ctr := 0
	nbrPos := [][]int{{i - 1, j}, {i - 1, j + 1}, {i, j + 1}, {i + 1, j + 1}, {i + 1, j}, {i + 1, j - 1}, {i, j - 1}, {i - 1, j - 1}}
	for _, pos := range nbrPos {
		x := pos[0]
		y := pos[1]
		if x >= 0 && x < m && y >= 0 && y < n && (board[x][y] == 1 || board[x][y] == 2) {
			ctr++
		}
	}
	return ctr
}

func nextStateOfCellNew(curr int, liveNbr int) int {
	// live to dead -> 2
	// dead to live -> 3

	if curr == 1 {
		if liveNbr < 2 {
			return 2
		} else if liveNbr < 4 {
			return 1
		} else {
			return 2
		}
	} else {
		if liveNbr == 3 {
			return 3
		} else {
			return 0
		}
	}
}
