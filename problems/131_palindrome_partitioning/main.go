package main

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

	var result [][]string

	var traverse func(int, []string)
	traverse = func(idx int, arr []string) {
		if idx == len(s) {
			result = append(result, arr)
			return
		}

		for i := idx; i < len(s); i++ {
			if dp[idx][i] {
				traverse(i+1, append(
					[]string{}, append(arr, s[idx:i+1])...,
				))
			}
		}
	}
	traverse(0, nil)

	return result
}
