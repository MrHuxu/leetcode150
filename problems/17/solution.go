package leetcode150

// code
var mapNumToLetters = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	var result []string
	for i := 0; i < len(digits); i++ {
		var next []string

		for _, b := range mapNumToLetters[digits[i]] {
			if len(result) == 0 {
				next = append(next, string([]byte{b}))
			}
			for _, pre := range result {
				next = append(next, string(append([]byte(pre), b)))
			}
		}
		result = next
	}
	return result
}
