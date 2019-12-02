package main

// code
func reverse(x int) int {
	return recursion(0, x)
}

func recursion(pre, num int) int {
	if pre > 2147483647 || pre < -2147483648 {
		return 0
	}
	if num == 0 {
		return pre
	}

	return recursion(pre*10+num%10, num/10)
}
