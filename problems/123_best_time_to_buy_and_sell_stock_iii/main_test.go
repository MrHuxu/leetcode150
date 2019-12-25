package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(6, maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	assert.Equal(4, maxProfit([]int{1, 2, 3, 4, 5}))
	assert.Equal(0, maxProfit([]int{7, 6, 4, 3, 1}))
	assert.Equal(0, maxProfit([]int{}))
}
