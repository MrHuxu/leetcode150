package leetcode150

// code
func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return s == ""
	}

	minLen := len(wordDict[0])
	maxLen := len(wordDict[0])
	exist := make(map[string]bool)
	for _, word := range wordDict {
		minLen = min(len(word), minLen)
		maxLen = max(len(word), maxLen)

		exist[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		if i < minLen {
			dp[i] = false
			continue
		}

		var found bool
		for j := minLen; j <= min(maxLen, i); j++ {
			if dp[i-j] && exist[s[i-j:i]] {
				found = true
				break
			}
		}
		dp[i] = found
	}

	return dp[len(s)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
