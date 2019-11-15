package main

import "sort"

// code
func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	var result [][]int

	var combine func(int, int, []int)
	combine = func(idx, left int, curr []int) {
		if left == 0 {
			result = append(result, curr)
			return
		}

		if left < 0 {
			return
		}

		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}

			combine(i+1, left-candidates[i], append([]int{}, append(curr, candidates[i])...))
		}
	}
	combine(0, target, []int{})

	return result
}
