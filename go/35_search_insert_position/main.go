package main

// code
func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		switch {
		case nums[mid] < target:
			left = mid + 1

		default:
			right = mid
		}
	}

	return (left + right) / 2
}
