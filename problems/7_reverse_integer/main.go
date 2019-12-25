package main

import "math"

// code
func reverse(x int) int {
	return helper(0, x)
}

func helper(pre, num int) int {
	if pre > math.MaxInt32 || pre < math.MinInt32 {
		return 0
	}
	if num == 0 {
		return pre
	}

	return helper(pre*10+num%10, num/10)
}
