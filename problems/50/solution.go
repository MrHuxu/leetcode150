package leetcode150

// code
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 1 || x == 0 {
		return x
	}
	if x == -1 {
		if n%2 == 0 {
			return 1
		}
		return -1
	}

	sign := 1
	if n < 0 {
		sign = -1
		n = -n
	}

	var pows = make(map[int]float64)
	pows[1] = x
	i := 2
	for ; i <= n; i += i {
		pows[i] = pows[i/2] * pows[i/2]
	}

	result := float64(1)
	for n != 0 {
		if n < i {
			i /= 2
			continue
		}

		result *= pows[i]
		n -= i
		i /= 2
	}

	if sign == -1 {
		result = 1 / result
	}
	return result
}
