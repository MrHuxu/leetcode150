package leetcode150

// code
func maxSubArray(nums []int) int {
	result := nums[0]
	pre := nums[0]

	for i := 1; i < len(nums); i++ {
		if pre < 0 {
			pre = nums[i]
		} else {
			pre = pre + nums[i]
		}

		if pre > result {
			result = pre
		}
	}

	return result
}
