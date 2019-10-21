package leetcode150

// code
func minDistance(word1, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 0; i <= len(word1); i++ {
		for j := 0; j <= len(word2); j++ {
			switch {
			case i == 0 && j == 0:
				dp[i][j] = 0

			case i == 0:
				dp[i][j] = j

			case j == 0:
				dp[i][j] = i

			default:
				if word1[i-1] == word2[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				}
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func min(a, b, c int) int {
	if b < a {
		a = b
	}

	if c < a {
		a = c
	}

	return a
}
