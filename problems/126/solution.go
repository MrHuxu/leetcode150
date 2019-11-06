package leetcode150

// code
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// list: map[ "*ot": []string{ "hot", "lot" } ]
	list := make(map[string][]string)
	for _, w := range wordList {
		for i := 0; i < len(w); i++ {
			bytes := []byte(w)
			bytes[i] = '*'
			list[string(bytes)] = append(list[string(bytes)], w)
		}
	}

	bfsResult, minDepth := bfs(beginWord, endWord, list)

	var result [][]string

	var dfs func(words []string, depth int)
	dfs = func(words []string, depth int) {
		if depth > minDepth {
			return
		}

		word := words[len(words)-1]

		if word == endWord {
			result = append(result, append([]string{}, words...))
			return
		}

		for _, nextWord := range bfsResult[word] {
			dfs(append(words, nextWord), depth+1)
		}
	}
	dfs([]string{beginWord}, 0)

	return result
}

func bfs(beginWord, endWord string, list map[string][]string) (map[string][]string, int) {
	result := make(map[string][]string)
	var depth int

	level := []string{beginWord}
	used := map[string]bool{beginWord: true}
	for len(level) != 0 {
		var found bool
		var nextLevel []string

		for _, w1 := range level {
			if w1 == endWord {
				found = true
				break
			}

			result[w1] = canChange(w1, list, used)
			nextLevel = append(nextLevel, result[w1]...)
		}
		if found {
			break
		}

		for _, w := range nextLevel {
			used[w] = true
		}
		depth++
		level = nextLevel
	}

	return result, depth
}

func canChange(word string, list map[string][]string, used map[string]bool) []string {
	var result []string

	for i := 0; i < len(word); i++ {
		bytes := []byte(word)
		bytes[i] = '*'

		for _, w := range list[string(bytes)] {
			if !used[w] {
				result = append(result, w)
			}
		}
	}

	return result
}
