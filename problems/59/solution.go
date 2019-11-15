package main

// code
func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	var num int
	for level := 0; level < (n+1)/2; level++ {
		for j := level; j < n-level; j++ {
			num++
			result[level][j] = num
		}

		for j := level + 1; j < n-level-1; j++ {
			num++
			result[j][n-level-1] = num
		}

		for j := n - level - 1; n-level-1 != level && j >= level; j-- {
			num++
			result[n-level-1][j] = num
		}

		for j := n - level - 2; n-level-1 != level && j > level; j-- {
			num++
			result[j][level] = num
		}
	}

	return result
}
