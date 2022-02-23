package main

// code
func subsets(nums []int) [][]int {
	var result [][]int

	var dfs func([]int, int)
	dfs = func(pre []int, idx int) {
		result = append(result, pre)

		for ; idx < len(nums); idx++ {
			dfs(append([]int{}, append(pre, nums[idx])...), idx+1)
		}
	}
	dfs([]int{}, 0)

	return result
}
