package leetcode150

// code
// divide and conquer
func largestRectangleArea1(heights []int) int {
	var result int

	var divideAndConquer func(int, int)
	divideAndConquer = func(left, right int) {
		if right < left {
			return
		}

		minIdx := left
		for i := left; i <= right; i++ {
			if heights[i] < heights[minIdx] {
				minIdx = i
			}
		}

		if heights[minIdx]*(right-left+1) > result {
			result = heights[minIdx] * (right - left + 1)
		}

		divideAndConquer(left, minIdx-1)
		divideAndConquer(minIdx+1, right)
	}
	divideAndConquer(0, len(heights)-1)

	return result
}

// stack
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
