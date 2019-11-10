package leetcode150

// code
func partition(s string) [][]string {
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

	return traverse(s, 0, dp)
}

func traverse(s string, i int, dp []map[int]bool) [][]string {
	var result [][]string

	for j := i; j < len(s); j++ {
		if dp[i][j] {
			left := s[i : j+1]
			rights := traverse(s, j+1, dp)

			for _, r := range rights {
				result = append(result, append([]string{left}, r...))
			}
			if len(rights) == 0 {
				result = append(result, []string{left})
			}
		}
	}

	return result
}
