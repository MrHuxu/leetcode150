package main

import (
	"sort"
)

// code
func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	var result [][]int

	var combine func(int, int, []int)
	combine = func(idx, sum int, arr []int) {
		if sum == target {
			result = append(result, arr)
			return
		}

		if sum > target {
			return
		}

		for i := idx; i < len(candidates); i++ {
			combine(i, sum+candidates[i], append([]int{}, append(arr, candidates[i])...))
		}
	}
	combine(0, 0, []int{})

	return result
}
