package leetcode150

// code
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var reversed int

	tmp := x
	for tmp > 0 {
		reversed = reversed*10 + tmp%10
		tmp /= 10
	}

	return reversed == x
}
