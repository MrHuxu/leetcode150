package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(7, maxProfit([]int{7, 1, 5, 3, 6, 4}))
	assert.Equal(4, maxProfit([]int{1, 2, 3, 4, 5}))
}
