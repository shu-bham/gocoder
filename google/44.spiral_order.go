package google

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	rowBegin, rowEnd := 0, m-1
	colBegin, colEnd := 0, n-1

	var ans []int
	dir := 0
	for rowBegin <= rowEnd && colBegin <= colEnd {
		switch dir {
		case 0: // process right
			for i := colBegin; i <= colEnd; i++ {
				ans = append(ans, matrix[rowBegin][i])
			}
			rowBegin++
		case 1: // process down
			for i := rowBegin; i <= rowEnd; i++ {
				ans = append(ans, matrix[i][colEnd])
			}
			colEnd--
		case 2: // process left
			for i := colEnd; i >= colBegin; i-- {
				ans = append(ans, matrix[rowEnd][i])
			}
			rowEnd--
		case 3: // process up
			for i := rowEnd; i >= rowBegin; i-- {
				ans = append(ans, matrix[i][colBegin])
			}
			colBegin++
		}
		dir = (dir + 1) % 4
	}

	return ans
}
