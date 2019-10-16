package leetcode150

// code
func removeDuplicates(nums []int) int {
	var preIndex int
	var moved bool

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[preIndex] {
			if !moved {
				moved = true
				preIndex++
				nums[preIndex] = nums[i]
			}
		} else {
			preIndex++
			moved = false
			nums[preIndex] = nums[i]
		}
	}

	return preIndex + 1
}
