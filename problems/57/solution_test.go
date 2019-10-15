package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insert(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{1, 5}, []int{6, 9},
	}, insert(
		[][]int{
			[]int{1, 3}, []int{6, 9},
		}, []int{2, 5},
	))

	assert.Equal([][]int{
		[]int{1, 2}, []int{3, 10}, []int{12, 16},
	}, insert([][]int{
		[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16},
	}, []int{4, 8}))

	assert.Equal([][]int{
		[]int{5, 7},
	}, insert([][]int{}, []int{5, 7}))
}
