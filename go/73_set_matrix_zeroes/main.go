package main

// code
func setZeroes(matrix [][]int) {
	var setFirstRow, setFirstCol bool
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			setFirstRow = true
			break
		}
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			setFirstCol = true
			break
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if setFirstRow {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}

	if setFirstCol {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

}
