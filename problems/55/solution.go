package leetcode150

// code
func canJump(nums []int) bool {
	var distance int
	for i, num := range nums {
		if distance >= len(nums)-1 {
			return true
		}

		if num == 0 && distance <= i {
			return false
		}

		distance = max(distance, num+i)
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
