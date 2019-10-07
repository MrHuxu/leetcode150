package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
