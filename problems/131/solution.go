package leetcode150

// code
func partition(s string) [][]string {
	dp := make([]map[int]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make(map[int]bool, len(s))
		dp[i][i] = true
	}

	return traverse(s, 0, len(s)-1, dp)
}

func traverse(s string, i, j int, dp []map[int]bool) [][]string {
	if i == j {
		return [][]string{
			[]string{s[i : j+1]},
		}
	}

	var result [][]string
	if validate(s, i, j, dp) {
		result = [][]string{
			[]string{s[i : j+1]},
		}
	}

	for k := i; k < j; k++ {
		if validate(s, i, k, dp) {
			left := s[i : k+1]
			rights := traverse(s, k+1, j, dp)

			for _, r := range rights {
				result = append(result, append([]string{left}, r...))
			}
		}
	}

	return result
}

func validate(s string, i, j int, dp []map[int]bool) bool {
	if i >= j {
		return true
	}

	result := s[i] == s[j]
	if result {
		if i+1 < len(s) {
			if _, ok := dp[i+1][j-1]; ok {
				result = result && dp[i+1][j-1]
			} else {
				result = result && validate(s, i+1, j-1, dp)
			}
		}
	}

	dp[i][j] = result
	return result
}
