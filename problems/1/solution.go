package leetcode150

func twoSum(nums []int, target int) []int {
	mapNumToPos := make(map[int]int)
	for i, v := range nums {
		if pos, ok := mapNumToPos[target-v]; ok {
			return []int{pos, i}
		}

		mapNumToPos[v] = i
	}
	return []int{}
}