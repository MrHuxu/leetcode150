package main

import "sort"

// code
func permuteUnique(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	return dfs(nums, make(map[int]bool))
}

func dfs(nums []int, used map[int]bool) [][]int {
	var ret [][]int

	for i, num := range nums {
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}

		used[i] = true
		for _, item := range dfs(nums, used) {
			ret = append(ret, append([]int{num}, item...))
		}
		used[i] = false
	}

	if len(ret) == 0 {
		return [][]int{{}}
	}

	return ret
}
