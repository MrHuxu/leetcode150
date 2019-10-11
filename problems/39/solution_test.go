package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{2, 2, 3},
		[]int{7},
	}, combinationSum([]int{2, 3, 6, 7}, 7))
	assert.Equal([][]int{
		[]int{2, 2, 2, 2},
		[]int{2, 3, 3},
		[]int{3, 5},
	}, combinationSum([]int{2, 3, 5}, 8))
	assert.Equal([][]int{
		[]int{3, 4, 4},
		[]int{3, 8},
		[]int{4, 7},
	}, combinationSum([]int{8, 7, 4, 3}, 11))
}
