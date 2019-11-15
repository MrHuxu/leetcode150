package main

// code
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}

	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}

			if i == 0 && j == 0 {
				dp[i][j] = 1
			}

			if i-1 >= 0 && i-1 < len(obstacleGrid) && obstacleGrid[i-1][j] != 1 {
				dp[i][j] += dp[i-1][j]
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] != 1 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
