package main

// code
func removeDuplicates(nums []int) int {
	var pos, pre int

	for i := 0; i < len(nums); i++ {
		if pos == 0 {
			pre = nums[0]
			pos++
			continue
		}

		num := nums[i]
		if num != pre {
			pre = num

			nums[pos] = num
			pos++
		}
	}

	return pos
}
