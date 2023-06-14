package main

// code
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2

		switch {
		case nums[mid] < nums[0]:
			right = mid - 1

		default:
			left = mid + 1
		}
	}
	rotateIdx := (left + right) / 2
	if target >= nums[0] {
		left, right = 0, rotateIdx
	} else {
		left, right = rotateIdx+1, len(nums)-1
	}

	for left <= right {
		mid := (left + right) / 2

		switch {
		case nums[mid] < target:
			left = mid + 1

		case nums[mid] > target:
			right = mid - 1

		default:
			return mid
		}
	}

	return -1
}
