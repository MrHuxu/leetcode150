package main

// code
func minCut(s string) int {
	dp := make([]map[int]bool, len(s))
	for i := range dp {
		dp[i] = make(map[int]bool)
	}

	for l := 1; l <= len(s); l++ {
		for i := 0; i+l <= len(s); i++ {
			j := i + l - 1

			switch l {
			case 1:
				dp[i][j] = true

			case 2:
				dp[i][j] = s[i] == s[j]

			default:
				dp[i][j] = dp[i+1][j-1] && (s[i] == s[j])
			}
		}
	}

	dp2 := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		dp2[i] = i

		for j := 0; j <= i; j++ {
			if dp[j][i] {
				if j == 0 {
					dp2[i] = 0
					break
				} else {
					dp2[i] = min(dp2[i], dp2[j-1]+1)
				}
			}
		}
	}

	return dp2[len(s)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
