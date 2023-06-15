package main

// code
// dp
func trap1(height []int) int {
	maxLeft, maxRight := make([]int, len(height)), make([]int, len(height))
	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(height[i-1], maxLeft[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = max(height[i+1], maxRight[i+1])
	}

	var result int
	for i := 1; i < len(height)-1; i++ {
		tmp := min(maxLeft[i], maxRight[i])
		if tmp > height[i] {
			result += tmp - height[i]
		}
	}

	return result
}

// backtrace
func trap2(height []int) int {
	if len(height) == 0 {
		return 0
	}

	var result int

	maxHeight := height[0]
	for i := 1; i < len(height); i++ {
		if height[i] > height[i-1] {
			tmp := min(height[i], maxHeight)

			for j := i - 1; j >= 0 && height[j] < tmp; j-- {
				result += tmp - height[j]
				height[j] = tmp
			}
		}

		maxHeight = max(maxHeight, height[i])
	}

	return result
}

// stack
func trap(height []int) int {
	var result int

	var stack []int
	for i := range height {
		for len(stack) != 0 && height[stack[len(stack)-1]] < height[i] {
			pre := height[stack[len(stack)-1]]

			stack = stack[0 : len(stack)-1]
			if len(stack) == 0 {
				break
			}

			distance := i - stack[len(stack)-1] - 1
			h := min(height[stack[len(stack)-1]], height[i])
			result += distance * (h - pre)
		}

		stack = append(stack, i)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
