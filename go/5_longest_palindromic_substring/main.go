package main

// code
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	var resultL, resultR int
	for l := 1; l <= len(s); l++ {
		for i := 0; i+l <= len(s); i++ {
			j := i + l - 1
			dp[i][j] = s[i] == s[j] && (l < 3 || dp[i+1][j-1])

			if dp[i][j] {
				resultL = i
				resultR = j
			}
		}
	}

	return s[resultL : resultR+1]
}
