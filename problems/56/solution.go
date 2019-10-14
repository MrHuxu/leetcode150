package leetcode150

// code
import "sort"

func merge(intervals [][]int) [][]int {
	var left, right []int
	for _, interval := range intervals {
		left = append(left, interval[0])
		right = append(right, interval[1])
	}
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	var result [][]int
	var stack []int
	for len(left) != 0 || len(right) != 0 {
		if len(right) == 0 || (len(left) != 0 && left[0] <= right[0]) {
			stack = append(stack, left[0])
			left = left[1:len(left)]
			continue
		}

		if len(left) == 0 || (len(right) != 0 && right[0] < left[0]) {
			preLeft := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]

			if len(stack) == 0 {
				result = append(result, []int{preLeft, right[0]})
			}
			right = right[1:len(right)]
		}
	}

	return result
}
