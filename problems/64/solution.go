package leetcode150

// code
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			dp[i][j] = grid[i][j]

			if i >= 1 && j >= 1 {
				dp[i][j] += min(dp[i-1][j], dp[i][j-1])
				continue
			}

			if i >= 1 {
				dp[i][j] += dp[i-1][j]
			}

			if j >= 1 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
