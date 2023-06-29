package main

import "sort"

// code
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var ret [][]int

	for _, interval := range intervals {
		if len(ret) == 0 || ret[len(ret)-1][1] < interval[0] {
			ret = append(ret, interval)
		} else {
			ret[len(ret)-1][1] = max(ret[len(ret)-1][1], interval[1])
		}
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
