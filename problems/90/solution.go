package main

import "sort"

// code
func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var result [][]int

	var traverse func(int, []int)
	traverse = func(idx int, curr []int) {
		result = append(result, curr)

		for i := idx; i < len(nums); i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}

			traverse(i+1, append([]int{}, append(curr, nums[i])...))
		}
	}
	traverse(0, []int{})

	return result
}
