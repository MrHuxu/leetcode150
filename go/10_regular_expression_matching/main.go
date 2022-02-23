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
				if p[i-1] == '*' {
					dp[i][j] = dp[i-2][0]
				} else {
					dp[i][j] = false
				}

			default:
				var matched bool

				if p[i-1] == s[j-1] || p[i-1] == '.' {
					matched = dp[i-1][j-1]
				}

				if !matched && p[i-1] == '*' {
					switch p[i-2] {
					case '.':
						for k := j; k >= 0; k-- {
							if dp[i-2][k] {
								matched = true
								break
							}
						}

					case s[j-1]:
						for k := j; k >= 0; k-- {
							if k == 0 {
								matched = dp[i-2][k]
								break
							}

							if dp[i-2][k] == true {
								matched = true
								break
							}

							if s[k-1] != s[j-1] {
								break
							}
						}

					default:
						matched = dp[i-2][j]
					}
				}

				dp[i][j] = matched
			}
		}
	}

	return dp[len(p)][len(s)]
}
