package leetcode150

// code
func firstMissingPositive(nums []int) int {
	var num int

	for len(nums) > 0 {
		for mid := len(nums) / 2; mid >= 1; mid-- {
			pos := mid - 1
			left := mid*2 - 1
			right := mid * 2

			if nums[left] < nums[pos] {
				tmp := nums[pos]
				nums[pos] = nums[left]
				nums[left] = tmp
			}

			if right < len(nums) && nums[right] < nums[pos] {
				tmp := nums[pos]
				nums[pos] = nums[right]
				nums[right] = tmp
			}
		}

		switch {
		case nums[0] <= 0:
			nums = nums[1:len(nums)]

		case nums[0] == num:
			nums = nums[1:len(nums)]

		case nums[0] == num+1:
			num++
			nums = nums[1:len(nums)]

		case nums[0] > num+1:
			return num + 1
		}
	}

	return num + 1
}
