package google

func CountUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for _, pos := range walls {
		x, y := pos[0], pos[1]
		ans[x][y] = -1 // -1 : wall
	}

	for _, pos := range guards {
		x, y := pos[0], pos[1]
		ans[x][y] = -2 // -2 : guard
	}

	var paint func(x int, y int, dir string)
	paint = func(x, y int, dir string) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if ans[x][y] == -1 || ans[x][y] == -2 {
			return
		}
		ans[x][y] = 1 // 1 : seen/ guarded
		var nextX, nextY int
		if dir == "right" {
			nextX = x
			nextY = y + 1
		} else if dir == "down" {
			nextX = x + 1
			nextY = y
		} else if dir == "left" {
			nextX = x
			nextY = y - 1
		} else if dir == "up" {
			nextX = x - 1
			nextY = y
		}
		paint(nextX, nextY, dir)
	}
	for _, pos := range guards {
		x, y := pos[0], pos[1]
		paint(x, y+1, "right")
		paint(x+1, y, "down")
		paint(x, y-1, "left")
		paint(x-1, y, "up")
	}

	ctr := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ans[i][j] == 0 {
				ctr++
			}
		}
	}
	return ctr
}
