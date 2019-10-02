package leetcode150

// code
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	var front, result int

	for i, b := range []byte(s) {
		if m[b] == 0 {
			if i-front+1 > result {
				result = i - front + 1
			}
		} else {
			for {
				if m[b] == 0 {
					break
				}

				m[s[front]]--
				front++
			}
		}
		m[b] = 1
	}

	return result
}
