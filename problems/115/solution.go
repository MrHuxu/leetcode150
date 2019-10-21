package leetcode150

// code
func numDistinct(S, T string) int {
	dp := make([][]int, len(S)+1)
	for i := range dp {
		dp[i] = make([]int, len(T)+1)
	}

	for i := 0; i <= len(S); i++ {
		for j := 0; j <= len(T); j++ {
			switch {
			case i == 0 && j == 0:
				dp[i][j] = 1

			case i == 0:
				dp[i][j] = 0

			case j == 0:
				dp[i][j] = 1

			default:
				dp[i][j] = dp[i-1][j]
				if S[i-1] == T[j-1] {
					dp[i][j] += dp[i-1][j-1]
				}
			}
		}
	}

	return dp[len(S)][len(T)]
}
