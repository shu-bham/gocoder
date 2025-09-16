package medium

/**
https://leetcode.com/problems/set-matrix-zeroes/description/
73. Set Matrix Zeroes

'Adobe', 'Amazon', 'Apple', 'Audible', 'Docusign', 'Expedia', 'Facebook', 'Microsoft', 'Oracle', 'Paypal', 'Salesforce', 'Tripadvisor'
*/

func SetZeroes(matrix [][]int) {

	m, n := len(matrix), len(matrix[0])
	rows := make([]bool, 200)
	cols := make([]bool, 200)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}

}
