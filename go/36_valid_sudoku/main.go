package main

// code
func isValidSudoku(board [][]byte) bool {
	rows := make([][]bool, 9)
	cols := make([][]bool, 9)
	rects := make([][]bool, 9)
	for i := range rows {
		rows[i] = make([]bool, 10)
		cols[i] = make([]bool, 10)
		rects[i] = make([]bool, 10)
	}

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				continue
			}

			if rects[i/3*3+j/3][b-48] {
				return false
			}
			rects[i/3*3+j/3][b-48] = true

			if rows[i][b-48] {
				return false
			}
			rows[i][b-48] = true

			if cols[j][b-48] {
				return false
			}
			cols[j][b-48] = true
		}
	}

	return true
}
