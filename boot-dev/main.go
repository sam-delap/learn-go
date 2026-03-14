package main

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = i * j
		}
	}

	return matrix
}

