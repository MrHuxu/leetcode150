package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	nums := []int{1, 2, 3}
	nextPermutation(nums)
	assert.Equal([]int{1, 3, 2}, nums)

	nums1 := []int{3, 2, 1}
	nextPermutation(nums1)
	assert.Equal([]int{1, 2, 3}, nums1)

	nums2 := []int{1, 1, 5}
	nextPermutation(nums2)
	assert.Equal([]int{1, 5, 1}, nums2)

	nums3 := []int{1, 3, 4, 2}
	nextPermutation(nums3)
	assert.Equal([]int{1, 4, 2, 3}, nums3)

	nums4 := []int{1, 4, 3, 2}
	nextPermutation(nums4)
	assert.Equal([]int{2, 1, 3, 4}, nums4)
}
