package leetcode150

import (
	"sort"
)

// code
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var result [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			switch {
			case nums[left]+nums[right]+nums[i] > 0:
				right--

			case nums[left]+nums[right]+nums[i] < 0:
				left++

			case nums[left]+nums[right]+nums[i] == 0:
				result = append(result, []int{nums[left], nums[i], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
				for left < right && nums[right-1] == nums[right] {
					right--
				}
				right--
			}
		}
	}

	return result
}
