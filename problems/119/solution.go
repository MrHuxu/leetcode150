package leetcode150

// code
func getRow(rowIndex int) []int {
	row := []int{1}

	for i := 1; i <= rowIndex; i++ {
		nextRow := make([]int, i+1)

		for j := 0; j < i+1; j++ {
			var val int

			if j > 0 {
				val = row[j-1]
			}
			if j < i {
				val += row[j]
			}

			nextRow[j] = val
		}

		row = nextRow
	}

	return row
}
