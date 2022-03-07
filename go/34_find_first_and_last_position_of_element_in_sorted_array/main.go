package main

// code
func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}

	if len(nums) == 0 {
		return ret
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
	if nums[(left+right)/2] == target {
		ret[0] = (left + right) / 2
	}

	left = 0
	right = len(nums) - 1
	for left <= right {
		mid := (left + right) / 2

		switch {
		case nums[mid] > target:
			right = mid - 1

		default:
			left = mid + 1
		}
	}
	if nums[(left+right)/2] == target {
		ret[1] = (left + right) / 2
	}

	return ret
}
