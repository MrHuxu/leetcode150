package leetcode150

// code
func isPalindrome(s string) bool {
	head := 0
	tail := len(s) - 1
	for head <= tail {
		ch1 := s[head]
		if ch1 >= 65 && ch1 <= 90 {
			ch1 += 32
		} else if !(ch1 >= 97 && ch1 <= 122) && !(ch1 >= 48 && ch1 <= 57) {
			head++
			continue
		}

		ch2 := s[tail]
		if ch2 >= 65 && ch2 <= 90 {
			ch2 += 32
		} else if !(ch2 >= 97 && ch2 <= 122) && !(ch2 >= 48 && ch2 <= 57) {
			tail--
			continue
		}

		if ch1 == ch2 {
			head++
			tail--
		} else {
			return false
		}
	}

	return true
}
