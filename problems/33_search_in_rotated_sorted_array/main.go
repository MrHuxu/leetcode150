package main

// code
func search(nums []int, target int) int {
	dividePos := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			dividePos = i
			break
		}
	}

	left := 0
	right := len(nums) - 1
	if dividePos == -1 {
		right = len(nums) - 1
	} else {
		if target > nums[0] {
			right = dividePos
		} else if target < nums[0] {
			left = dividePos + 1
		} else {
			return 0
		}
	}

	for left <= right {
		mid := (left + right) / 2

		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
