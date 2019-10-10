package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement(t *testing.T) {
	assert := assert.New(t)

	arr := []int{3, 2, 2, 3}
	assert.Equal(2, removeElement(arr, 3))
	assert.Equal(2, arr[0])
	assert.Equal(2, arr[1])

	arr = []int{0, 1, 2, 2, 3, 0, 4, 2}
	assert.Equal(5, removeElement(arr, 2))
	assert.Equal(0, arr[0])
	assert.Equal(1, arr[1])
	assert.Equal(3, arr[2])
	assert.Equal(0, arr[3])
	assert.Equal(4, arr[4])
}
