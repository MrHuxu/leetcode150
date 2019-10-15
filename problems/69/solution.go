package leetcode150

// code
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	left := 1
	right := x / 2
	for left <= right {
		mid := (left + right) / 2

		pow := mid * mid
		if pow == x {
			return mid
		} else if pow > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}
