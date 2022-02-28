package main

// code
func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}

	wordCnt := make(map[string]int)
	for _, word := range words {
		wordCnt[word]++
	}
	wordLen := len(words[0])

	var dfs func(int, int, map[string]int) bool
	dfs = func(idx, left int, wordCnt map[string]int) bool {
		if left == 0 {
			return true
		}

		subStr := s[idx:(idx + wordLen)]
		if wordCnt[subStr] <= 0 {
			return false
		}

		wordCnt[subStr]--
		found := dfs(idx+wordLen, left-1, wordCnt)
		wordCnt[subStr]++

		return found
	}

	var result []int
	for i := 0; i <= len(s)-(wordLen*len(words)); i++ {
		if dfs(i, len(words), wordCnt) {
			result = append(result, i)
		}
	}

	return result
}
