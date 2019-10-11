package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{2, 2, 3},
		[]int{7},
	}, test([]int{2, 3, 6, 7}, 7))
	assert.Equal([][]int{
		[]int{2, 2, 2, 2},
		[]int{2, 3, 3},
		[]int{3, 5},
	}, test([]int{2, 3, 5}, 8))
}
