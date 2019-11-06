package leetcode150

// code
func totalNQueens(n int) int {
	var result int

	used := make([]int, n)
	var dfs func(board [][]byte, i int)
	dfs = func(board [][]byte, i int) {
		if i == n {
			result++
			return
		}

		for j := 0; j < n; j++ {
			if validate(i, j, used) {
				board[i][j] = 'Q'
				used[i] = j

				dfs(board, i+1)

				board[i][j] = 0
			}
		}
	}

	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
	}

	dfs(board, 0)

	return result
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
