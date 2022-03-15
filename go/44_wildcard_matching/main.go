package main

// code
func isMatch(s, p string) bool {
	dp := make([][]bool, len(p)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}

	for i := 0; i <= len(p); i++ {
		for j := 0; j <= len(s); j++ {
			switch {
			case i == 0 && j == 0:
				dp[i][j] = true

			case i == 0:
				dp[i][j] = false

			case j == 0:
				dp[i][j] = p[i-1] == '*' && dp[i-1][0]

			default:
				switch p[i-1] {
				case '*':
					dp[i][j] = dp[i-1][j-1] || dp[i][j-1] || dp[i-1][j]

				case '?':
					dp[i][j] = dp[i-1][j-1]

				default:
					dp[i][j] = dp[i-1][j-1] && p[i-1] == s[j-1]
				}
			}
		}
	}

	return dp[len(p)][len(s)]
}
