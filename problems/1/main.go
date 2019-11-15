package main

// code
func twoSum(nums []int, target int) []int {
	mapNumToPos := make(map[int]int, len(nums))
	for i, v := range nums {
		if pos, ok := mapNumToPos[target-v]; ok {
			return []int{pos, i}
		}

		mapNumToPos[v] = i
	}
	return []int{}
}
