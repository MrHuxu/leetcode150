package main

// code
func removeDuplicates(nums []int) int {
	var preIndex int
	var same bool

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[preIndex] {
			if !same {
				same = true
				preIndex++
				nums[preIndex] = nums[i]
			}
		} else {
			preIndex++
			same = false
			nums[preIndex] = nums[i]
		}
	}

	return preIndex + 1
}
