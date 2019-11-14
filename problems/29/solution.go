package leetcode150

import "math"

// code
func divide(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	}

	var result int
	sign := 1

	if dividend > 0 {
		sign = opposite(sign)
		dividend = opposite(dividend)
	}
	if divisor > 0 {
		sign = opposite(sign)
		divisor = opposite(divisor)
	}

	for dividend <= divisor {
		for i := uint(0); dividend <= (divisor << i); i++ {
			dividend -= (divisor << i)
			result += int(1 << i)
		}
	}

	if sign != 1 {
		result = opposite(result)
	}

	if result > math.MaxInt32 {
		result = math.MaxInt32
	}

	return result
}

func opposite(i int) int {
	return ^i + 1
}
