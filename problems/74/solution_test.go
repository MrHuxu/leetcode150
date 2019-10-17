package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchMatrix(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, searchMatrix([][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}, 3))

	assert.Equal(false, searchMatrix([][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}, 13))

	assert.Equal(false, searchMatrix([][]int{
		[]int{1},
	}, 0))
}
