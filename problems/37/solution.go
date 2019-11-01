package leetcode150

// code
func solveSudoku(board [][]byte) {
	dfs(board, 0, 0)
}

func dfs(board [][]byte, i, j int) bool {
	if i == 9 {
		return true
	}

	if j == 9 {
		return dfs(board, i+1, 0)
	}

	if board[i][j] != '.' {
		return dfs(board, i, j+1)
	}

	for ch := byte('1'); ch <= '9'; ch++ {
		if validate(board, i, j, ch) {
			board[i][j] = ch

			if dfs(board, i, j+1) {
				return true
			}

			board[i][j] = '.'
		}
	}

	return false
}

func validate(board [][]byte, i, j int, ch byte) bool {
	for _, col := range board[i] {
		if col == ch {
			return false
		}
	}

	for _, row := range board {
		if row[j] == ch {
			return false
		}
	}

	for m := i / 3 * 3; m < (i/3+1)*3; m++ {
		for n := j / 3 * 3; n < (j/3+1)*3; n++ {
			if board[m][n] == ch {
				return false
			}
		}
	}

	return true
}
