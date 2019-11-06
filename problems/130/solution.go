package leetcode150

// code
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if board[i][j] == 'X' {
			return
		}

		visited[i][j] = true
		if i-1 >= 0 && !visited[i-1][j] {
			dfs(i-1, j)
		}
		if j-1 >= 0 && !visited[i][j-1] {
			dfs(i, j-1)
		}
		if i+1 < len(board) && !visited[i+1][j] {
			dfs(i+1, j)
		}
		if j+1 < len(board[0]) && !visited[i][j+1] {
			dfs(i, j+1)
		}
	}

	for i := 0; i < len(board[0]); i++ {
		dfs(0, i)
	}
	for i := 1; i < len(board)-1; i++ {
		dfs(i, 0)
		dfs(i, len(board[0])-1)
	}
	if len(board) > 1 {
		for i := 0; i < len(board[0]); i++ {
			dfs(len(board)-1, i)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
