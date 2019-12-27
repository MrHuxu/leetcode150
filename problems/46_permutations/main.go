package main

// code
func permute(nums []int) [][]int {
	var result [][]int

	used := make([]bool, len(nums))
	var dfs func([]int, int)
	dfs = func(arr []int, depth int) {
		if depth == len(nums) {
			result = append(result, arr)
		}

		for i := 0; i < len(nums); i++ {
			if !used[i] {
				used[i] = true
				dfs(append([]int{}, append(arr, nums[i])...), depth+1)
				used[i] = false
			}
		}
	}
	dfs(nil, 0)

	return result
}
