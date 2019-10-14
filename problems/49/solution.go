package leetcode150

// code
type node struct {
	index int
	next  map[int]*node
}

func groupAnagrams(strs []string) [][]string {
	var result [][]string

	root := &node{
		index: -1,
		next:  make(map[int]*node),
	}
	for _, str := range strs {
		times := make([]int, 26)
		for _, b := range []byte(str) {
			times[b-97]++
		}

		tmp := root
		for i, time := range times {
			for time != 0 {
				if _, ok := tmp.next[i]; !ok {
					tmp.next[i] = &node{
						index: -1,
						next:  make(map[int]*node),
					}
				}
				tmp = tmp.next[i]

				time--
			}
		}

		if tmp.index == -1 {
			result = append(result, []string{})
			tmp.index = len(result) - 1
		}
		result[tmp.index] = append(result[tmp.index], str)
	}

	return result
}
