package main

// code
func exist(board [][]byte, word string) bool {
	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[i]))
	}

	var dfs func(int, int, int) bool
	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}

		if i-1 >= 0 && board[i-1][j] == word[k] && !used[i-1][j] {
			used[i-1][j] = true
			if dfs(i-1, j, k+1) {
				return true
			}
			used[i-1][j] = false
		}
		if i+1 < len(board) && board[i+1][j] == word[k] && !used[i+1][j] {
			used[i+1][j] = true
			if dfs(i+1, j, k+1) {
				return true
			}
			used[i+1][j] = false
		}
		if j-1 >= 0 && board[i][j-1] == word[k] && !used[i][j-1] {
			used[i][j-1] = true
			if dfs(i, j-1, k+1) {
				return true
			}
			used[i][j-1] = false
		}
		if j+1 < len(board[0]) && board[i][j+1] == word[k] && !used[i][j+1] {
			used[i][j+1] = true
			if dfs(i, j+1, k+1) {
				return true
			}
			used[i][j+1] = false
		}

		return false
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				used[i][j] = true
				if dfs(i, j, 1) {
					return true
				}
				used[i][j] = false
			}
		}
	}

	return false
}
