package main

// code
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		num := nums[i]
		if num >= 1 && num <= n && nums[num-1] != num {
			nums[i], nums[num-1] = nums[num-1], nums[i]
			i--
		}
	}

	for i, num := range nums {
		if i+1 != num {
			return i + 1
		}
	}

	return n + 1
}
