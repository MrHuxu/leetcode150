package leetcode150

// code
func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] ^= nums[i-1]
	}

	return nums[len(nums)-1]
}
