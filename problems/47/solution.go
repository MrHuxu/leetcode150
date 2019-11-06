package leetcode150

import "sort"

// code
func permuteUnique(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	result := [][]int{nums}

	for {
		next := append([]int{}, result[len(result)-1]...)
		adjusted := nextPermutation(next)

		if adjusted {
			result = append(result, next)
		} else {
			break
		}
	}

	return result
}

func nextPermutation(nums []int) (adjusted bool) {
	for i := len(nums) - 2; i >= 0; i-- {
		var hasBigger bool
		var minBiggerIdx int

		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				if hasBigger {
					if nums[j] < nums[minBiggerIdx] {
						minBiggerIdx = j
					}
				} else {
					minBiggerIdx = j
					hasBigger = true
				}
			}
		}

		if hasBigger {
			tmp := nums[minBiggerIdx]
			nums[minBiggerIdx] = nums[i]
			nums[i] = tmp

			subNums := nums[i+1 : len(nums)]
			sort.Slice(subNums, func(a, b int) bool {
				return subNums[a] < subNums[b]
			})

			adjusted = true
			return
		}
	}

	return
}
