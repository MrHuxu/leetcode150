package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insert(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		{1, 5}, {6, 9},
	}, insert(
		[][]int{
			{1, 3}, {6, 9},
		}, []int{2, 5},
	))

	assert.Equal([][]int{
		{1, 2}, {3, 10}, {12, 16},
	}, insert([][]int{
		{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	}, []int{4, 8}))

	assert.Equal([][]int{
		{5, 7},
	}, insert([][]int{}, []int{5, 7}))

	assert.Equal([][]int{
		{0, 5},
	}, insert([][]int{{1, 5}}, []int{0, 3}))

	assert.Equal([][]int{
		{0, 0},
		{1, 5},
	}, insert([][]int{{1, 5}}, []int{0, 0}))
}
