package leetcode150

// code
func test(nums []int) int {
	num := 1

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
			num++
			nums = nums[1:len(nums)]

		case nums[0] != num:
			return num
		}
	}

	return num
}
