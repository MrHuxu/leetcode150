package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_spiralOrder(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{
		1, 2, 3, 6, 9, 8, 7, 4, 5,
	}, spiralOrder([][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}))

	assert.Equal([]int{6, 9, 7}, spiralOrder([][]int{
		[]int{6, 9, 7},
	}))
}
