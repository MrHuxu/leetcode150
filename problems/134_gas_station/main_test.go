package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canCompleteCircuit(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, canCompleteCircuit(
		[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2},
	))

	assert.Equal(-1, canCompleteCircuit(
		[]int{2, 3, 4}, []int{3, 4, 3},
	))
}
