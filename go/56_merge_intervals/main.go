package main

import "sort"

// code
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int

	for i := 0; i < len(intervals); i++ {
		left := intervals[i][0]
		right := intervals[i][1]

		for j := i + 1; j < len(intervals); j++ {
			if intervals[j][0] <= right {
				if intervals[j][1] > right {
					right = intervals[j][1]
				}

				i = j
			} else {
				break
			}
		}
		result = append(result, []int{left, right})
	}

	return result
}
