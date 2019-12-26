package main

// code
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	pre := nums[0]
	pos := 1

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num != pre {
			pre = num

			nums[pos] = num
			pos++
		}
	}

	return pos
}
