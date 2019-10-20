package leetcode150

// code
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	var resultLeft, resultRight int

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i+(resultRight-resultLeft); j-- {
			if s[i] == s[j] && isPalindrome(i, j, s) {
				resultLeft = i
				resultRight = j

				break
			}
		}
	}

	return s[resultLeft : resultRight+1]
}

func isPalindrome(i, j int, s string) bool {
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
