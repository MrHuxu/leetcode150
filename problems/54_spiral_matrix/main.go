package main

// code
func spiralOrder(matrix [][]int) []int {
	var result []int

	for level := 0; len(matrix) > 0 && len(result) < len(matrix)*len(matrix[0]); level++ {
		for i := level; i < len(matrix[0])-level; i++ {
			result = append(result, matrix[level][i])
		}

		for i := level + 1; i < len(matrix)-level-1; i++ {
			result = append(result, matrix[i][len(matrix[0])-level-1])
		}

		for i := len(matrix[0]) - level - 1; level != len(matrix)-level-1 && i >= level; i-- {
			result = append(result, matrix[len(matrix)-level-1][i])
		}

		for i := len(matrix) - level - 2; level != len(matrix[0])-level-1 && i > level; i-- {
			result = append(result, matrix[i][level])
		}
	}

	return result
}
