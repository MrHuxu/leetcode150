package main

// code
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	var targetRow []int

	rowLen := len(matrix[0])
	left, right := 0, len(matrix)-1
	for left <= right {
		mid := (left + right) / 2

		if matrix[mid][0] <= target && target <= matrix[mid][rowLen-1] {
			targetRow = matrix[mid]
			break
		} else if matrix[mid][0] > target {
			right = mid - 1
		} else if matrix[mid][rowLen-1] < target {
			left = mid + 1
		}
	}

	if targetRow == nil {
		return false
	}

	left = 0
	right = rowLen
	for left <= right {
		mid := (left + right) / 2

		if targetRow[mid] == target {
			return true
		} else if targetRow[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
