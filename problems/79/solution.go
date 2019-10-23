package leetcode150

// code
func exist(board [][]byte, word string) bool {
	var dfs func([][]bool, int, int, int) bool
	dfs = func(used [][]bool, i, j, k int) bool {
		if k == len(word) {
			return true
		}

		if i-1 >= 0 && board[i-1][j] == word[k] && !used[i-1][j] {
			copied := copy(used)
			copied[i-1][j] = true
			if dfs(copied, i-1, j, k+1) {
				return true
			}
		}
		if i+1 < len(board) && board[i+1][j] == word[k] && !used[i+1][j] {
			copied := copy(used)
			copied[i+1][j] = true
			if dfs(copied, i+1, j, k+1) {
				return true
			}
		}
		if j-1 >= 0 && board[i][j-1] == word[k] && !used[i][j-1] {
			copied := copy(used)
			copied[i][j-1] = true
			if dfs(copied, i, j-1, k+1) {
				return true
			}
		}
		if j+1 < len(board[0]) && board[i][j+1] == word[k] && !used[i][j+1] {
			copied := copy(used)
			copied[i][j+1] = true
			if dfs(copied, i, j+1, k+1) {
				return true
			}
		}

		return false
	}

	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[i]))
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				copied := copy(used)
				copied[i][j] = true
				if dfs(copied, i, j, 1) {
					return true
				}
			}
		}
	}

	return false
}

func copy(used [][]bool) [][]bool {
	result := make([][]bool, len(used))
	for i := range result {
		result[i] = make([]bool, len(used[i]))
		for j := range used[i] {
			result[i][j] = used[i][j]
		}
	}
	return result
}
