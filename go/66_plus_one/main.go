package main

// code
func plusOne(digits []int) []int {
	carry := 1

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry

		carry = digits[i] / 10
		digits[i] %= 10
	}

	if carry == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}
