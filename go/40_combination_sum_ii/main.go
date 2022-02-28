package main

import "sort"

// code
func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	var result [][]int

	var combine func(int, int, []int)
	combine = func(idx, sum int, curr []int) {
		if sum == target {
			result = append(result, curr)
			return
		}

		if sum > target {
			return
		}

		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}

			combine(i+1, sum+candidates[i], append([]int{}, append(curr, candidates[i])...))
		}
	}
	combine(0, 0, []int{})

	return result
}
