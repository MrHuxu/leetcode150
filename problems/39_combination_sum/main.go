package main

// code
import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
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
			combine(i, left-candidates[i], append([]int{}, append(curr, candidates[i])...))
		}
	}
	combine(0, target, []int{})

	return result
}
