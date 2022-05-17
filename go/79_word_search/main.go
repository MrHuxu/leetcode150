package main

// code
func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	for x := range board {
		for y := range board[x] {
			if board[x][y] != word[0] {
				continue
			}

			visited[x][y] = true
			if dfs(board, word, x, y, 1, visited) {
				return true
			}
			visited[x][y] = false
		}
	}

	return false
}

func dfs(board [][]byte, word string, x, y, idx int, visited [][]bool) bool {
	if idx == len(word) {
		return true
	}

	if x > 0 && board[x-1][y] == word[idx] && !visited[x-1][y] {
		visited[x-1][y] = true
		if dfs(board, word, x-1, y, idx+1, visited) {
			return true
		}
		visited[x-1][y] = false
	}
	if x < len(board)-1 && board[x+1][y] == word[idx] && !visited[x+1][y] {
		visited[x+1][y] = true
		if dfs(board, word, x+1, y, idx+1, visited) {
			return true
		}
		visited[x+1][y] = false
	}
	if y > 0 && board[x][y-1] == word[idx] && !visited[x][y-1] {
		visited[x][y-1] = true
		if dfs(board, word, x, y-1, idx+1, visited) {
			return true
		}
		visited[x][y-1] = false
	}
	if y < len(board[0])-1 && board[x][y+1] == word[idx] && !visited[x][y+1] {
		visited[x][y+1] = true
		if dfs(board, word, x, y+1, idx+1, visited) {
			return true
		}
		visited[x][y+1] = false
	}
	return false
}
