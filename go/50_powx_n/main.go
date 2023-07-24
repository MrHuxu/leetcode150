package main

import "math"

// code
func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}

	sym := n > 0
	n = int(math.Abs(float64(n)))

	pow2Vals := map[int]float64{}
	pow2, val := 1, x
	for pow2 <= n {
		pow2Vals[pow2] = val
		pow2 *= 2
		val *= val
	}

	ret := float64(1)
	for n != 0 {
		if pow2 > n {
			pow2 /= 2
			continue
		}

		n -= pow2
		ret *= pow2Vals[pow2]
		pow2 /= 2
	}

	if !sym {
		return 1 / ret
	}
	return ret
}
