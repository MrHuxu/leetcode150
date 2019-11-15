package main

// code
func solveNQueens(n int) [][]string {
	var result [][]string

	used := make([]int, n)
	var dfs func(board [][]byte, i int)
	dfs = func(board [][]byte, i int) {
		if i == n {
			result = append(result, copyBoard(board))
			return
		}

		for j := 0; j < n; j++ {
			if validate(i, j, used) {
				board[i][j] = 'Q'
				used[i] = j

				dfs(board, i+1)

				board[i][j] = '.'
			}
		}
	}

	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	dfs(board, 0)

	return result
}

func copyBoard(board [][]byte) []string {
	copied := make([]string, len(board))

	for i := range copied {
		copied[i] = string(board[i])
	}

	return copied
}

func validate(i, j int, used []int) bool {
	for k := 0; k < i; k++ {
		if used[k] == j {
			return false
		}
	}

	for k := 0; k < i; k++ {
		if k-i == used[k]-j || k-i == j-used[k] {
			return false
		}
	}

	return true
}
