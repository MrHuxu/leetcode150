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

	min := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		tmp := i

		for j := 0; j <= i; j++ {
			if dp[j][i] {
				if j == 0 {
					tmp = 0
					break
				} else {
					if min[j-1]+1 < tmp {
						tmp = min[j-1] + 1
					}
				}
			}
		}

		min[i] = tmp
	}

	return min[len(s)-1]
}
