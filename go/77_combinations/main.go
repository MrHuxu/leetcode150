package main

// code
func combine(n, k int) [][]int {
	var result [][]int

	var dfs func([]int, int)
	dfs = func(pre []int, curr int) {
		if len(pre) == k {
			result = append(result, pre)
			return
		}

		for ; curr <= n-k+1+len(pre); curr++ {
			dfs(append([]int{}, append(pre, curr)...), curr+1)
		}
	}
	dfs(nil, 1)

	return result
}
