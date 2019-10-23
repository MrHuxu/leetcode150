package leetcode150

// code
func grayCode(n int) []int {
	result := []int{0}

	for i := 0; i < n; i++ {
		for j := len(result) -1; j >= 0; j-- {
			result = append(result, result[j] + (1<<uint(i)))
		}
	}

	return result
}