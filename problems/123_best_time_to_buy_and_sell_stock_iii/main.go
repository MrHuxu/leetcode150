package main

// code
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	for i := 1; i <= 2; i++ {
		minVal := prices[0]

		for j := 1; j <= len(prices); j++ {
			minVal = min(prices[j-1]-dp[j][i-1], minVal)
			dp[j][i] = max(dp[j-1][i], prices[j-1]-minVal)
		}
	}

	return dp[len(prices)][2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
