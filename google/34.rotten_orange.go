package google

import "fmt"

type point struct {
	r, c int
}

func OrangesRotting(grid [][]int) int {
	queue := make([]point, 0)
	m, n := len(grid), len(grid[0])

	freshCount := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				freshCount++
			} else if grid[i][j] == 2 {
				queue = append(queue, point{i, j})
			}
		}
	}

	fmt.Println(m, n, queue)

	diff := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	time := 0
	for len(queue) > 0 && freshCount > 0 {
		// fmt.Println(time, freshCount, grid)
		k := len(queue)
		time++
		// process all elements at this depth
		for i := 0; i < k; i++ {
			front := queue[0]
			queue = queue[1:]

			for _, d := range diff {
				x := front.r + d[0]
				y := front.c + d[1]
				// fmt.Println("processing:", x, y)
				if x >= 0 && x < m && y >= 0 && y < n {
					if grid[x][y] == 1 {
						// fmt.Println("processing fresh fruit")
						freshCount--
						grid[x][y] = 2
						queue = append(queue, point{x, y})
					}
				}
			}
		}
	}

	if freshCount > 0 {
		return -1
	}
	return time

}
