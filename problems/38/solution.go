package leetcode150

import "strconv"

// code
func countAndSay(n int) string {
	str := "1"
	for i := 2; i <= n; i++ {
		var next string

		letter := str[0]
		length := 1
		for i := 1; i < len(str); i++ {
			if str[i] == letter {
				length++
			} else {
				next += strconv.Itoa(length)
				next += string([]byte{letter})
				letter = str[i]
				length = 1
			}
		}

		next += strconv.Itoa(length)
		next += string([]byte{letter})
		str = next
	}

	return str
}
