package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2.0, findMedianSortedArrays(
		[]int{1, 3}, []int{2},
	))
	assert.Equal(2.5, findMedianSortedArrays(
		[]int{1, 2}, []int{3, 4},
	))
	assert.Equal(0.0, findMedianSortedArrays(
		[]int{0, 0}, []int{0, 0},
	))
	assert.Equal(2.0, findMedianSortedArrays(
		[]int{2}, []int{},
	))
}
