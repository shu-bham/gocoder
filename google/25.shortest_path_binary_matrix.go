package google

func ShortestPathBinaryMatrix(grid [][]int) int {
	N := len(grid)
	if grid[N-1][N-1] == 1 || grid[0][0] == 1 {
		return -1
	}

	type tuple struct {
		x, y, dist int
	}

	var queue []tuple
	queue = append(queue, tuple{0, 0, 1})
	grid[0][0] = 1

	all := [][]int{{-1, -1}, {-1, 0}, {-1, +1}, {0, -1}, {0, +1}, {+1, -1}, {+1, 0}, {+1, +1}}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		x := front.x
		y := front.y

		if x == N-1 && y == N-1 {
			return front.dist
		}
		for _, pos := range all {
			x0, y0 := x+pos[0], y+pos[1]
			if x0 >= 0 && x0 < N && y0 >= 0 && y0 < N && grid[x0][y0] == 0 {
				grid[x0][y0] = 1
				queue = append(queue, tuple{x0, y0, front.dist + 1})
			}
		}
	}
	return -1
}
