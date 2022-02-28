package main

// code
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			switch {
			case i == 0 && j == 0:
				dp[i][j] = true

			case i == 0:
				dp[i][j] = s3[0:j] == s2[0:j]

			case j == 0:
				dp[i][j] = s3[0:i] == s1[0:i]

			default:
				dp[i][j] = (s2[j-1] == s3[i+j-1] && dp[i][j-1]) || (s1[i-1] == s3[i+j-1] && dp[i-1][j])
			}
		}
	}

	return dp[len(s1)][len(s2)]
}
