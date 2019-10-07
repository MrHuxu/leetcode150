package leetcode150

func maxArea(height []int) int {
	var result int

	var left int
	right := len(height) - 1
	for left < right {
		if height[left] < height[right] {
			area := height[left] * (right - left)
			if area > result {
				result = area
			}
			left++
		} else {
			area := height[right] * (right - left)
			if area > result {
				result = area
			}
			right--
		}
	}

	return result
}
