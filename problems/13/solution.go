package leetcode150

var mapSymToNum = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var result int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'V', 'X':
			for j := i - 1; j >= 0 && s[j] == 'I'; j-- {
				result -= mapSymToNum['I'] * 2
			}

		case 'L', 'C':
			for j := i - 1; j >= 0 && s[j] == 'X'; j-- {
				result -= mapSymToNum['X'] * 2
			}

		case 'D', 'M':
			for j := i - 1; j >= 0 && s[j] == 'C'; j-- {
				result -= mapSymToNum['C'] * 2
			}
		}
		result += mapSymToNum[s[i]]
	}
	return result
}
