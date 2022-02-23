package main

// code
func removeElement(nums []int, val int) int {
	var pos int

	for i := range nums {
		num := nums[i]

		if num != val {
			nums[pos] = num
			pos++
		}
	}

	return pos
}
