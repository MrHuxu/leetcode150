package leetcode150

import "fmt"

// code
func longestValidParentheses(s string) int {
	var result int

	dp := make([]int, len(s)+1)
	for i := 1; i < len(s); i++ {
		fmt.Println(dp)
		if s[i] == ')' {
			if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
				dp[i] = dp[i-1] + 2

				if i-2-dp[i-1] >= 0 && dp[i-2-dp[i-1]] > 0 {
					dp[i] += dp[i-2-dp[i-1]]
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
