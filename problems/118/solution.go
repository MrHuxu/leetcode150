package leetcode150

// code
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	rows := [][]int{[]int{1}}

	for i := 2; i <= numRows; i++ {
		rows = append(rows, make([]int, i))

		for j := 0; j < i; j++ {
			var val int

			if j > 0 {
				val = rows[i-2][j-1]
			}
			if j < i-1 {
				val += rows[i-2][j]
			}

			rows[i-1][j] = val
		}
	}

	return rows
}
