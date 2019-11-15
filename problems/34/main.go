package main

// code
func searchRange(nums []int, target int) []int {
	var found bool
	var pos int

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			found = true
			pos = mid
			break
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if !found {
		return []int{-1, -1}
	}

	first := pos
	for first > 0 && nums[first-1] == nums[first] {
		first--
	}
	last := pos
	for last < len(nums)-1 && nums[last+1] == nums[last] {
		last++
	}

	return []int{first, last}
}
