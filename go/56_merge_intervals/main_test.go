package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		{1, 6},
		{8, 10},
		{15, 18},
	}, merge([][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}))
}
