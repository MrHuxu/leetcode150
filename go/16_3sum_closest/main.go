package main

import (
	"math"
	"sort"
)

// code
func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	minSub := math.MaxInt32
	var result int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			switch {
			case sum > target:
				if sum-target < minSub {
					minSub = sum - target
					result = sum
				}
				right--

			case sum < target:
				if target-sum < minSub {
					minSub = target - sum
					result = sum
				}
				left++

			default:
				return target
			}
		}
	}

	return result
}
