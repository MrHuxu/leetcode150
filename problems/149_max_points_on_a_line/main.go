package main

import "strconv"

// code
func maxPoints(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}

	var result int

	for i := 0; i < len(points); i++ {
		m := make(map[string]int)

		var samePos, sameCol int
		tmp := 1
		for j := i + 1; j < len(points); j++ {
			horizontalDiff := points[j][0] - points[i][0]
			verticalDiff := points[j][1] - points[i][1]

			if horizontalDiff == 0 {
				if verticalDiff == 0 {
					samePos++
				} else {
					sameCol++
					tmp = max(tmp, sameCol+1)
				}
			} else {
				divisor := gcd(horizontalDiff, verticalDiff)
				key := strconv.Itoa(verticalDiff/divisor) + "@" + strconv.Itoa(horizontalDiff/divisor)
				m[key]++

				tmp = max(tmp, m[key]+1)
			}
		}

		result = max(result, tmp+samePos)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for b != 0 {
		tmp := a % b
		a = b
		b = tmp
	}
	return a
}
