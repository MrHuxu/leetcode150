package main

// code
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// list: map[ "*ot": []string{ "hot", "lot" } ]
	list := make(map[string][]string)
	for _, w := range wordList {
		for i := 0; i < len(w); i++ {
			bytes := []byte(w)
			bytes[i] = '*'
			list[string(bytes)] = append(list[string(bytes)], w)
		}
	}

	return bfs(beginWord, endWord, list)
}

func bfs(beginWord, endWord string, changeList map[string][]string) int {
	result := make(map[string][]string)
	depth := 1

	level := []string{beginWord}
	used := map[string]bool{beginWord: true}
	for len(level) != 0 {
		var found bool
		var nextLevel []string

		for _, w1 := range level {
			if w1 == endWord {
				return depth
			}

			result[w1] = canChange(w1, changeList, used)
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

	return 0
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
