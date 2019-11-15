package main

// code
func subsets(nums []int) [][]int {
	return dfs([]int{}, nums, 0)
}

func dfs(pre, nums []int, curr int) [][]int {
	result := [][]int{pre}

	for ; curr < len(nums); curr++ {
		result = append(
			result, dfs(append([]int{}, append(pre, nums[curr])...), nums, curr+1)...,
		)
	}
	return result
}
