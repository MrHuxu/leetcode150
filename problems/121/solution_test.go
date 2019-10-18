package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
	assert.Equal(0, maxProfit([]int{7, 6, 4, 3, 1}))
}
