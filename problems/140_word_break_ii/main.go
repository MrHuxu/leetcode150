package main

// code
func wordBreak(s string, wordDict []string) []string {
	if len(wordDict) == 0 {
		return nil
	}

	minLen := len(wordDict[0])
	maxLen := len(wordDict[0])
	exist := make(map[string]bool)
	for _, word := range wordDict {
		exist[word] = true

		l := len(word)
		if l < minLen {
			minLen = l
		}
		if l > maxLen {
			maxLen = l
		}
	}

	used := make(map[int][]string)
	var dfs func(int) []string
	dfs = func(idx int) []string {
		var result []string

		if result, ok := used[idx]; ok {
			return result
		}

		for j := minLen; j <= maxLen && idx+j <= len(s); j++ {
			subS := s[idx : idx+j]

			if exist[subS] {
				if idx+j == len(s) {
					result = append(result, subS)
				} else {
					left := dfs(idx + j)
					if len(left) > 0 {
						for _, l := range left {
							result = append(result, subS+" "+l)
						}
					}
				}
			}
		}

		used[idx] = result
		return result
	}

	return dfs(0)
}
