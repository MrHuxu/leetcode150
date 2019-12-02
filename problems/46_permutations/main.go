package main

// code
func permute(nums []int) [][]int {
	return dfs(nums, []int{}, make([]bool, len(nums)), 0)
}

func dfs(origin, curr []int, used []bool, depth int) [][]int {
	if depth == len(origin) {
		return [][]int{curr}
	}

	var result [][]int
	for i := 0; i < len(origin); i++ {
		if !used[i] {
			newUsed := append([]bool{}, used...)
			newUsed[i] = true
			result = append(result, dfs(
				origin, append([]int{}, append(curr, origin[i])...), newUsed, depth+1,
			)...)
		}
	}
	return result
}
