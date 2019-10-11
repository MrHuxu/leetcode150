package leetcode150

// code
func test(candidates []int, target int) [][]int {
	return combine(candidates, []int{}, 0, 0, target)
}

func combine(candidates, arr []int, pos, sum, target int) [][]int {
	if sum == target {
		return [][]int{arr}
	} else if sum > target {
		return nil
	}

	var result [][]int
	for j := pos; j < len(candidates); j++ {
		candidate := candidates[j]

		if sum+candidate > target {
			break
		} else {
			result = append(result, combine(
				candidates, append(arr, candidate), j, sum+candidate, target,
			)...)
		}
	}
	return result
}
