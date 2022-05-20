package main

// code
func reverse(x int) int {
	return helper(0, int32(x))
}

func helper(pre, num int32) int {
	if num == 0 {
		return int(pre)
	}

	if (num > 0 && (pre*10+num%10)/10 < pre) || (num < 0 && (pre*10+num%10)/10 > pre) {
		return 0
	}

	return helper(pre*10+num%10, num/10)
}
