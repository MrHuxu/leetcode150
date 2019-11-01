package leetcode150

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

			rect := rects[rectangleIdx(i, j)]
			if rect[b-48] {
				return false
			}
			rect[b-48] = true

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

func rectangleIdx(i, j int) int {
	switch {
	case i < 3:
		switch {
		case j < 3:
			return 0

		case j >= 3 && j < 6:
			return 1

		case j >= 6:
			return 2
		}

	case i >= 3 && i < 6:
		switch {
		case j < 3:
			return 3

		case j >= 3 && j < 6:
			return 4

		case j >= 6:
			return 5
		}

	case i >= 6:
		switch {
		case j < 3:
			return 6

		case j >= 3 && j < 6:
			return 7

		case j >= 6:
			return 8
		}

	}
	return 0
}
