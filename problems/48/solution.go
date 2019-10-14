package leetcode150

// code
func rotate(matrix [][]int) {
	for i := 0; (i+1)*2 <= len(matrix); i++ {
		for j := i; j < len(matrix)-i-1; j++ {
			tmp1 := matrix[j][len(matrix)-1-i]
			matrix[j][len(matrix)-1-i] = matrix[i][j]

			tmp2 := matrix[len(matrix)-1-i][len(matrix)-1-j]
			matrix[len(matrix)-1-i][len(matrix)-1-j] = tmp1

			tmp3 := matrix[len(matrix)-1-j][i]
			matrix[len(matrix)-1-j][i] = tmp2

			matrix[i][j] = tmp3
		}
	}
}
