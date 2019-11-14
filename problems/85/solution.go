package leetcode150

// code
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	var result int

	heights := make([]int, len(matrix[0]))
	for _, row := range matrix {
		for i, col := range row {
			switch col {
			case '0':
				heights[i] = 0

			case '1':
				heights[i] = heights[i] + 1
			}
		}
		result = max(result, largestRectangleArea(heights))
	}

	return result
}

func largestRectangleArea(heights []int) int {
	var result int

	var stack []int
	for i := 0; i < len(heights); {
		if len(stack) == 0 || heights[i] >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
			continue
		}

		height := heights[stack[len(stack)-1]]
		stack = stack[0 : len(stack)-1]

		var left int
		if len(stack) == 0 {
			left = -1
		} else {
			left = stack[len(stack)-1]
		}
		right := i

		area := (right - left - 1) * height
		result = max(result, area)
	}

	for len(stack) != 0 {
		height := heights[stack[len(stack)-1]]
		stack = stack[0 : len(stack)-1]

		var left int
		if len(stack) == 0 {
			left = -1
		} else {
			left = stack[len(stack)-1]
		}
		right := len(heights)

		area := (right - left - 1) * height
		result = max(result, area)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
