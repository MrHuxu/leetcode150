package leetcode150

// code
func combine(n, k int) [][]int {
	return dfs([]int{}, 1, n, k)
}

func dfs(pre []int, curr, n, k int) [][]int {
	if len(pre) == k {
		return [][]int{pre}
	}

	var result [][]int
	for ; curr <= n-k+1+len(pre); curr++ {
		result = append(
			result, dfs(append([]int{}, append(pre, curr)...), curr+1, n, k)...,
		)
	}
	return result
}
