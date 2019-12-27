package main

// code
func totalNQueens(n int) int {
	var result int

	used := make([]int, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			result++
			return
		}

		for j := 0; j < n; j++ {
			if validate(i, j, used) {
				used[i] = j
				dfs(i+1)
				used[i] = 0
			}
		}
	}
	dfs(0)

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
