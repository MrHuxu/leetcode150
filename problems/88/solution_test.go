package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	arr := []int{1, 2, 3, 0, 0, 0}
	merge(arr, 3, []int{2, 5, 6}, 3)
	assert.Equal([]int{1, 2, 2, 3, 5, 6}, arr)
}
