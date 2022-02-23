package main

import "sort"

// code
func nextPermutation(nums []int) {
	var adjusted bool

	for i := len(nums) - 2; i >= 0; i-- {
		var hasBigger bool
		var minBiggerIdx int

		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				if hasBigger {
					if nums[j] < nums[minBiggerIdx] {
						minBiggerIdx = j
					}
				} else {
					minBiggerIdx = j
					hasBigger = true
				}
			}
		}

		if hasBigger {
			tmp := nums[minBiggerIdx]
			nums[minBiggerIdx] = nums[i]
			nums[i] = tmp

			subNums := nums[i+1 : len(nums)]
			sort.Slice(subNums, func(a, b int) bool {
				return subNums[a] < subNums[b]
			})

			adjusted = true
			break
		}
	}

	if !adjusted {
		sort.Slice(nums, func(a, b int) bool {
			return nums[a] < nums[b]
		})
	}
}
