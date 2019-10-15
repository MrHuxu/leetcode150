package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minPathSum(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(7, minPathSum([][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}))
}
