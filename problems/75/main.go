package main

// code
func sortColors(nums []int) {
	pos0 := -1
	pos1 := -1
	pos2 := -1

	for _, num := range nums {
		switch num {
		case 0:
			pos2++
			nums[pos2] = 2
			pos1++
			nums[pos1] = 1
			pos0++
			nums[pos0] = 0

		case 1:
			pos2++
			nums[pos2] = 2
			pos1++
			nums[pos1] = 1

		case 2:
			pos2++
			nums[pos2] = 2
		}
	}
}
