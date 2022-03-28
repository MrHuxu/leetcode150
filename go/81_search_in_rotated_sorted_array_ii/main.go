package main

// code
func search(nums []int, target int) bool {
	if target == nums[0] {
		return true
	}

	rotateIdx := findRotateIdx(nums, 0, len(nums)-1)
	var left, right int
	if target > nums[0] {
		left = 0
		right = *rotateIdx
	} else {
		left = *rotateIdx + 1
		right = len(nums) - 1
	}

	for left <= right {
		mid := (left + right) / 2

		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}

func findRotateIdx(nums []int, left, right int) *int {
	for left <= right {
		mid := (left + right) / 2
		if mid == len(nums)-1 || nums[mid] > nums[mid+1] {
			return &mid
		}

		if nums[mid] < nums[0] {
			right = mid - 1
			continue
		}
		if nums[mid] > nums[0] {
			left = mid + 1
			continue
		}

		if idx := findRotateIdx(nums, left, mid-1); idx != nil {
			return idx
		}
		return findRotateIdx(nums, mid+1, right)
	}
	return nil
}
