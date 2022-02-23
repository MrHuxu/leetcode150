package main

// code
func longestValidParentheses(s string) int {
	var result int

	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
				dp[i] = dp[i-1] + 2

				if i-dp[i] >= 0 && dp[i-dp[i]] > 0 {
					dp[i] += dp[i-dp[i]]
				}

				result = max(result, dp[i])
			}
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
